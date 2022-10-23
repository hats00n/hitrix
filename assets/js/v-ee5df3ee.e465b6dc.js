"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[7152],{1866:(e,s,n)=>{n.r(s),n.d(s,{data:()=>a});const a={key:"v-ee5df3ee",path:"/guide/services/html2pdf.html",title:"HTML2PDF service",lang:"en-US",frontmatter:{},excerpt:"",headers:[],filePathRelative:"guide/services/html2pdf.md",git:{updatedTime:1659017697e3,contributors:[{name:"Krasimir Ivanov",email:"krasimir.ivanov@coretrix.com",commits:1}]}}},2378:(e,s,n)=>{n.r(s),n.d(s,{default:()=>c});const a=(0,n(6252).uE)('<h1 id="html2pdf-service" tabindex="-1"><a class="header-anchor" href="#html2pdf-service" aria-hidden="true">#</a> HTML2PDF service</h1><p>HTML2PDF service provides a generating pdf function from html code using Chrome headless.</p><p>First you need these in your app config:</p><div class="language-yaml ext-yml line-numbers-mode"><pre class="language-yaml"><code><span class="token key atrule">chrome_headless</span><span class="token punctuation">:</span>\n  <span class="token key atrule">web_socket_url</span><span class="token punctuation">:</span> ENV<span class="token punctuation">[</span>CHROME_HEADLESS_WEB_SOCKET_URL<span class="token punctuation">]</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br></div></div><p>Register the service into your <code>main.go</code> file:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>registry<span class="token punctuation">.</span><span class="token function">ServiceProviderHTML2PDF</span><span class="token punctuation">(</span><span class="token punctuation">)</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><p>Access the service:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>service<span class="token punctuation">.</span><span class="token function">DI</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">HTML2PDFService</span><span class="token punctuation">(</span><span class="token punctuation">)</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><p>Using <code>HtmlToPdf()</code> function to generate PDF from html:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>pdfBytes <span class="token operator">:=</span> html2pdfService<span class="token punctuation">.</span><span class="token function">HtmlToPdf</span><span class="token punctuation">(</span><span class="token string">&quot;&lt;html&gt;&lt;p&gt;Hi!&lt;/p&gt;&lt;/html&gt;&quot;</span><span class="token punctuation">)</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><p>Recommended docker file for Chrome headless:</p><div class="language-text ext-text line-numbers-mode"><pre class="language-text"><code>https://hub.docker.com/r/chromedp/headless-shell/\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div>',12),t={},c=(0,n(3744).Z)(t,[["render",function(e,s){return a}]])},3744:(e,s)=>{s.Z=(e,s)=>{for(const[n,a]of s)e[n]=a;return e}}}]);