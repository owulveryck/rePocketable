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

img {
	text-align: center !important;
	aspect-ratio: auto 1441 / 922;
	box-sizing: border-box;
	font: inherit;
	vertical-align: baseline;
	margin: 0 auto;
	padding: 0;
	border: 0;
	max-width: 100%;
	object-fit: contain;
	margin-bottom: 8px;
	width: auto!important;
	height: auto!important;
}

figure {
	color: var(#3d3b49);
	text-align: left;
	box-sizing: border-box;
	font: inherit;
	vertical-align: baseline;
	page-break-inside: avoid;
	-webkit-margin-before: 0;
	-webkit-margin-after: 0;
	display: block;
	border: none !important;
	margin: 32px 0 !important;
	padding: 0 !important;
}

figure div {
	box-sizing: border-box;
	font: inherit;
	vertical-align: baseline;
	line-height: 1;
	font-family: serif;
	text-align: center !important;
	page-break-inside: avoid;
	background-color: inherit;
	-webkit-margin-before: 0;
	-webkit-margin-after: 0;
	border: none !important;
	margin: 32px 0 !important;
	padding: 0 !important;
}

table {
	border-collapse: collapse;
	border-spacing: 0;
}

ol, ul, li, {
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
	margin: 0 0 24px 0 !important;
	font-size: 1.5em;
	font-weight: bold;
	page-break-before: always;
	/*get squished otherwise on ADE */
	line-height: 1.5 !important;
}

h3 {
	box-sizing: border-box;
	font: inherit;
	vertical-align: baseline;
	hyphens: none;
	text-align: left;
	page-break-after: avoid !important;
	-webkit-margin-before: 0;
	-webkit-margin-after: 0;
	background-color: transparent;
	color: #3d3b49 !important;
	padding: 0 !important;
	border: none !important;
	font-family: "Noto Serif", serif !important;
	font-style: normal;
	word-wrap: break-word;
	font-size: 1.125em !important;
	font-weight: bold !important;
	line-height: 1.5 !important;
	margin: 0 0 24px 0 !important;
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

div > h6 {
	font-family: Guardian Text Sans\ 2, sans-serif !important;
	font-size: 0.688em !important;
	font-weight: bold !important;
	line-height: 1 !important;
	text-align: left;
	margin: 14px 0 7px 0 !important;
	text-transform: uppercase;
}

p {
	margin: 0 0 24px 0 !important;
	line-height: 1.55 !important;
}

p[data-type="attribution"] {
-webkit-text-size-adjust: 100%;
color: var(#3d3b49);
quotes: none;
box-sizing: border-box;
border: 0;
font: inherit;
vertical-align: baseline;
width: 90%;
padding: 0;
-webkit-margin-before: 0;
-webkit-margin-after: 0;
font-family: "Noto serif", serif !important;
font-size: 1em;
font-weight: 400;
line-height: 1.75 !important;
hyphens: auto;
margin: 0 0 24px 0 !important;
margin-bottom: 0 !important;
text-align: right;
font-style: normal !important;
margin-top: 0 !important;
}

dl {
	text-size-adjust: 100%;
	color: #3d3b49;
	text-align: left;
	box-sizing: border-box;
	border: 0px;
	font: inherit;
	vertical-align: baseline;
	padding: 0px;
	margin-block: 0px;
	list-style: none;
	margin: 24px 0px 24px 24px !important;
}

dt {
	text-size-adjust: 100%;
	color: #3d3b49;
	text-align: left;
	list-style: none;
	box-sizing: border-box;
	border: 0px;
	font: inherit;
	vertical-align: baseline;
	font-style: italic;
	margin: 0px;
	margin-block: 0px;
	font-weight: 400;
	hyphens: auto;
	font-size: 1em !important;
	line-height: 1.75 !important;
	padding: 0px !important;	
}

dd {
	text-size-adjust: 100%;
	color: #3d3b49;
	list-style: none;
	box-sizing: border-box;
	border: 0px;
	font: inherit;
	vertical-align: baseline;
	margin: 10px 0px 0.25em 1.5em !important;
	line-height: 1.65em !important;
	text-align: left;
	padding: 0px;
	margin-block: 0px;
	font-size: 90%;
}

div[data-type="note"] {
    border: none;
    border-top: 1px solid #8B889A;
    border-bottom: 1px solid #8B889A;
    padding: 0 !important;
    margin: 40px 0 !important;
    color: #363636;
    line-height: 1.25 !important;
    border-radius: 0;
}

div[data-type="note"] p {
	font-size: 90%;
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
	color: var(#3d3b49);
	text-align: left;
	box-sizing: border-box;
	border: 0;
	font: inherit;
	vertical-align: baseline;
	quotes: none;
	page-break-inside: avoid;
	padding: 0;
	-webkit-margin-before: 0;
	-webkit-margin-after: 0;
	font-family: "Noto serif", serif !important;
	font-size: 1em;
	font-weight: 400;
	line-height: 1.75 !important;
	hyphens: auto;
	margin: 0 24px 24px !important;
	font-style: italic;
}
/*
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

figure h6 {
	box-sizing: border-box;
	font: inherit;
	vertical-align: baseline;
	text-align: center;
	text-transform: none !important;
	letter-spacing: normal !important;
	page-break-before: avoid;
	-webkit-margin-before: 0;
	-webkit-margin-after: 0;
	background-color: transparent;
	color: #3d3b49 !important;
	padding: 0 !important;
	border: none !important;
	font-style: normal;
	word-wrap: break-word;
	font-family: "Noto serif", serif !important;
	font-size: 0.750em !important;
	font-weight: 400 !important;
	line-height: 1.125 !important;
	hyphens: auto;
	margin: 0 0 24px 0 !important;
}

blockquote p {
	display: inline;
}
*/

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
