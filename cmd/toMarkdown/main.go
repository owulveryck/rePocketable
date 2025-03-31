package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	nethttp "net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/owulveryck/rePocketable/internal/http"
	"github.com/owulveryck/rePocketable/internal/markdown"
	"github.com/owulveryck/rePocketable/internal/pocket"
)

var DB map[string]pocket.Item

type headers map[string][]string

func (h headers) String() string {
	var b strings.Builder
	for k, v := range h {
		fmt.Fprintf(&b, "%v: %v|", k, v)
	}
	return b.String()
}

func (h headers) Set(v string) error {
	elements := strings.SplitN(v, ":", 2)
	if len(elements) != 2 {
		return errors.New("bad header passed")
	}
	h[elements[0]] = append(h[elements[0]], elements[1])
	return nil
}

func main() {
	// Check if any arguments are provided
	if len(os.Args) <= 1 || (len(os.Args) > 1 && os.Args[1] == "-h") {
		// No arguments, start MCP server
		startMCPServer()
		return
	}

	// Arguments provided, run in CLI mode
	runCLIMode()
}

func startMCPServer() {
	// Create MCP server
	s := server.NewMCPServer(
		"ToMarkdown 📄",
		"1.0.0",
	)

	// Add ToMarkdown tool
	tool := mcp.NewTool("ToMarkdown",
		mcp.WithDescription("Converts a web page to markdown. Provide a URL and receive the content converted to markdown format. The tool handles downloading the content, extracting the main text, and formatting it as clean markdown."),
		mcp.WithString("url",
			mcp.Required(),
			mcp.Description("The URL of the web page to convert to markdown"),
		),
	)

	// Add tool handler
	s.AddTool(tool, toMarkdownHandler)

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func toMarkdownHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	url, ok := request.Params.Arguments["url"].(string)
	if !ok {
		return nil, errors.New("url must be a string")
	}

	// Initialize HTTP client
	downloader, err := http.NewDownloader(nethttp.Header{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize HTTP client: %v", err)
	}

	// Create document
	item := pocket.Item{
		ResolvedURL: url,
		GivenURL:    url,
	}
	
	doc := markdown.NewDocument(item)
	doc.Client = downloader.HTTPClient
	
	// Fill document
	err = doc.Fill(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot fill document: %v", err)
	}
	
	// Instead of writing to a file, capture the output
	var buf bytes.Buffer
	err = doc.WriteTo(&buf)
	if err != nil {
		return nil, fmt.Errorf("cannot convert document to markdown: %v", err)
	}
	
	// Return markdown content
	return mcp.NewToolResultText(buf.String()), nil
}

func runCLIMode() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	var headersFlag headers
	headersFlag = make(map[string][]string)
	flag.Var(&headersFlag, "H", "header")
	
	if usage() {
		return
	}

	downloader, err := http.NewDownloader(nethttp.Header(headersFlag))
	if err != nil {
		log.Fatal(err)
	}

	item := pocket.Item{
		ResolvedURL: os.Args[len(os.Args)-1],
		GivenURL:    os.Args[len(os.Args)-1],
	}
	
	doc := markdown.NewDocument(item)
	doc.Client = downloader.HTTPClient
	
	err = doc.Fill(ctx)
	if err != nil {
		log.Println("Cannot fill document: ", err)
		return
	}
	
	outputFilename := fmt.Sprintf("%v.md", filepath.Base(os.Args[len(os.Args)-1]))
	log.Println("writing output: ", outputFilename)
	
	err = doc.Write(outputFilename)
	if err != nil {
		log.Fatal("Cannot write document: ", err)
	}
}