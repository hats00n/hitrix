"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[315],{213:(n,s,a)=>{a.r(s),a.d(s,{data:()=>e});const e={key:"v-d33a8d88",path:"/guide/services/mailjet.html",title:"Mail Mailjet service",lang:"en-US",frontmatter:{},excerpt:"",headers:[],filePathRelative:"guide/services/mailjet.md",git:{updatedTime:1656076836e3,contributors:[{name:"Masih Yeganeh",email:"goodboy.php@gmail.com",commits:4}]}}},471:(n,s,a)=>{a.r(s),a.d(s,{default:()=>p});const e=(0,a(6252).uE)('<h1 id="mail-mailjet-service" tabindex="-1"><a class="header-anchor" href="#mail-mailjet-service" aria-hidden="true">#</a> Mail Mailjet service</h1><p>This service is used for sending transactional emails using Mailjet</p><p>Register the service into your <code>main.go</code> file:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>registry<span class="token punctuation">.</span><span class="token function">ServiceProviderMail</span><span class="token punctuation">(</span>mail<span class="token punctuation">.</span>NewMailjet<span class="token punctuation">)</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><p>and you should register the entity <code>MailTrackerEntity</code> into the ORM Also, you should put your credentials and other configs in your config file</p><div class="language-yaml ext-yml line-numbers-mode"><pre class="language-yaml"><code><span class="token key atrule">mail</span><span class="token punctuation">:</span>\n  <span class="token key atrule">mailjet</span><span class="token punctuation">:</span>\n    <span class="token key atrule">api_key_public</span><span class="token punctuation">:</span> <span class="token punctuation">...</span>\n    <span class="token key atrule">api_key_private</span><span class="token punctuation">:</span> <span class="token punctuation">...</span>\n    <span class="token key atrule">default_from_email</span><span class="token punctuation">:</span> test@coretrix.tv\n    <span class="token key atrule">from_name</span><span class="token punctuation">:</span> coretrix.com\n    <span class="token key atrule">sandbox_mode</span><span class="token punctuation">:</span> <span class="token boolean important">false</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br><span class="line-number">5</span><br><span class="line-number">6</span><br><span class="line-number">7</span><br></div></div><p>Access the service:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>service<span class="token punctuation">.</span><span class="token function">DI</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">.</span><span class="token function">Mail</span><span class="token punctuation">(</span><span class="token punctuation">)</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br></div></div><p>Some functions this service provide are:</p><div class="language-go ext-go line-numbers-mode"><pre class="language-go"><code>\t<span class="token function">SendTemplate</span><span class="token punctuation">(</span>ormService <span class="token operator">*</span>beeorm<span class="token punctuation">.</span>Engine<span class="token punctuation">,</span> message <span class="token operator">*</span>TemplateMessage<span class="token punctuation">)</span> <span class="token builtin">error</span>\n\t<span class="token function">SendTemplateAsync</span><span class="token punctuation">(</span>ormService <span class="token operator">*</span>beeorm<span class="token punctuation">.</span>Engine<span class="token punctuation">,</span> message <span class="token operator">*</span>TemplateMessage<span class="token punctuation">)</span> <span class="token builtin">error</span>\n\t<span class="token function">SendTemplateWithAttachments</span><span class="token punctuation">(</span>ormService <span class="token operator">*</span>beeorm<span class="token punctuation">.</span>Engine<span class="token punctuation">,</span> message <span class="token operator">*</span>TemplateAttachmentMessage<span class="token punctuation">)</span> <span class="token builtin">error</span>\n\t<span class="token function">SendTemplateWithAttachmentsAsync</span><span class="token punctuation">(</span>ormService <span class="token operator">*</span>beeorm<span class="token punctuation">.</span>Engine<span class="token punctuation">,</span> message <span class="token operator">*</span>TemplateAttachmentMessage<span class="token punctuation">)</span> <span class="token builtin">error</span>\n</code></pre><div class="line-numbers"><span class="line-number">1</span><br><span class="line-number">2</span><br><span class="line-number">3</span><br><span class="line-number">4</span><br></div></div>',10),t={},p=(0,a(3744).Z)(t,[["render",function(n,s){return e}]])},3744:(n,s)=>{s.Z=(n,s)=>{for(const[a,e]of s)n[a]=e;return n}}}]);