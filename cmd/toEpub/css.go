package main

import "bytes"

var css = bytes.NewBufferString(`
/*===Reset code to prevent cross-reader strangeness===*/

html, body, div, span, applet, object, iframe, h1, h2, h3, h4, h5, h6, p, blockquote, pre, a, abbr, acronym, address, big, cite, code, del, dfn, em, img, ins, kbd, q, s, samp, small, strike, strong, sub, sup, tt, var, b, u, i, center, fieldset, form, label, legend, table, caption, tbody, tfoot, thead, tr, th, td, article, aside, canvas, details, embed, figure, figcaption, footer, header, hgroup, menu, nav, output, ruby, section, summary, time, mark, audio, video {
	margin: 0;
	padding: 0;
	border: 0;
	font-size: 100%;
	vertical-align: baseline;
}

table {
	border-collapse: collapse;
	border-spacing: 0;
}

ol, ul, li, dl, dt, dd {
	margin: 0;
	padding: 0;
	border: 0;
	font-size: 100%;
	vertical-align: baseline;
}

h1 {
	text-indent: 0;
	text-align: center;
	margin: 100px 0 0 0;
	font-size: 2.0em;
	font-weight: bold;
	page-break-before: always;
	line-height: 150%;
	/*gets squished otherwise on ADE */
}

h2 {
	text-indent: 0;
	text-align: center;
	margin: 50px 0 0 0;
	font-size: 1.5em;
	font-weight: bold;
	page-break-before: always;
	line-height: 135%;
	/*get squished otherwise on ADE */
}

h3 {
	text-indent: 0;
	text-align: left;
	font-size: 1.4em;
	font-weight: bold;
}

h4 {
	text-indent: 0;
	text-align: left;
	font-size: 1.2em;
	font-weight: bold;
}

h5 {
	text-indent: 0;
	text-align: left;
	font-size: 1.1em;
	font-weight: bold;
}

h6 {
	text-indent: 0;
	text-align: left;
	font-size: 1.0em;
	font-weight: bold;
}

/* Hyphen and pagination Fixer */

/* Note: Do not try on the Kindle, it does not recognize the hyphens property */

h1, h2, h3, h4, h5, h6 {
	-webkit-hyphens: none !important;
	hyphens: none;
	page-break-after: avoid;
	page-break-inside: avoid;
}

/*==LISTS==*/

ul {
	margin: 1em 0 0 2em;
	text-align: left;
}

ol {
	margin: 1em 0 0 2em;
	text-align: left;
}

span.i {
	font-style: italic;
}

span.b {
	font-weight: bold;
}

span.u {
	text-decoration: underline;
}

span.st {
	text-decoration: line-through;
}

/*==in-line combinations==*/

/* Using something like <span class="i b">... may seem okay, but it causes problems on some eReaders */

span.ib {
	font-style: italic;
	font-weight: bold;
}

span.iu {
	font-style: italic;
	text-decoration: underline;
}

span.bu {
	font-weight: bold;
	text-decoration: underline;
}

span.ibu {
	font-style: italic;
	font-weight: bold;
	text-decoration: underline;
}

div.pullquote {
	margin: 2em 2em 0 2em;
	text-align: left;
}

div.pullquote p {
	font-weight: bold;
	font-style: italic;
}

div.pullquote hr {
	width: 100%;
	margin: 0;
	height: 3px;
	color: #2E8DE0;
	background-color: #2E8DE0;
	border: 0;
}

/* Code in text */

p>code, li>code, dd>code, td>code {
	background: #ffeff0;
	word-wrap: break-word;
	padding: 0.1rem 0.1rem 0.1rem;
}

/* quotes etc... */
blockquote {
	background: #f9f9f9;
	border-left: 10px solid #ccc;
	margin: 1.5em 10px;
	padding: 0.5em 10px;
	quotes: "\201C""\201D""\2018""\2019";
}

blockquote:before {
	color: #ccc;
	content: open-quote;
	font-size: 4em;
	line-height: 0.1em;
	margin-right: 0.25em;
	vertical-align: -0.4em;
}

blockquote p {
	display: inline;
}

figcaption {
	--reach-tabs: 1;
	--reach-menu-button: 1;
	text-rendering: optimizeLegibility;
	-webkit-font-smoothing: antialiased;
	word-break: break-word;
	word-wrap: break-word;
	box-sizing: inherit;
	font-weight: 400;
	font-family: sohne, "Helvetica Neue", Helvetica, Arial, sans-serif;
	font-size: 14px;
	line-height: 20px;
	color: rgba(117, 117, 117, 1);
	margin-right: auto;
	max-width: 728px;
	margin-top: 10px;
	margin-left: auto;
	text-align: center;
}`)
