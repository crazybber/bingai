if(!self.define){let e,s={};const i=(i,n)=>(i=new URL(i+".js",n).href,s[i]||new Promise((s=>{if("document"in self){const e=document.createElement("script");e.src=i,e.onload=s,document.head.appendChild(e)}else e=i,importScripts(i),s()})).then((()=>{let e=s[i];if(!e)throw new Error(`Module ${i} didn’t register its module`);return e})));self.define=(n,c)=>{const a=e||("document"in self?document.currentScript.src:"")||location.href;if(s[a])return;let r={};const o=e=>i(e,a),d={module:{uri:a},exports:r,require:o};s[a]=Promise.all(n.map((e=>d[e]||o(e)))).then((e=>(c(...e),r)))}}define(["./workbox-118fddf1"],(function(e){"use strict";self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"assets/index-0e8bf8a6.css",revision:null},{url:"assets/index-2e609ae8.js",revision:null},{url:"assets/index-6956b245.js",revision:null},{url:"assets/index-9d6f4de0.css",revision:null},{url:"assets/setting-c6ca7b14.svg",revision:null},{url:"compose.html",revision:"8b03b32410e58485c18402934281fede"},{url:"favicon.ico",revision:"1272c70e1b86b8956598a0349d2f193c"},{url:"img/compose.svg",revision:"4242b76bb8f4da0baf7a75edab0c6754"},{url:"img/logo.svg",revision:"1da58864f14c1a8c28f8587d6dcbc5d0"},{url:"img/pwa/logo-192.png",revision:"be40443731d9d4ead5e9b1f1a6070135"},{url:"img/pwa/logo-512.png",revision:"1217f1c90acb9f231e3135fa44af7efc"},{url:"index.html",revision:"dd4db415146bc050257b530bde295618"},{url:"js/bing/chat/amd.js",revision:"8d773dc8f2e78b9d29e990aed7821774"},{url:"js/bing/chat/config.js",revision:"0e83c78343088c1cae24e2779e1f4c4d"},{url:"js/bing/chat/core.js",revision:"8c11521fd9f049b6ac91e5ad415c2db1"},{url:"js/bing/chat/global.js",revision:"43fd640c3a3021437b9b2955e8f2d1ad"},{url:"js/bing/chat/lib.js",revision:"1a0f8f43cc025b7b5995e885fed1a3e6"},{url:"registerSW.js",revision:"bf6c2f29aef95e09b1f72cf59f427a55"},{url:"./img/pwa/logo-192.png",revision:"be40443731d9d4ead5e9b1f1a6070135"},{url:"./img/pwa/logo-512.png",revision:"1217f1c90acb9f231e3135fa44af7efc"},{url:"manifest.webmanifest",revision:"ae4ef030ae5d2d4894669fd82aac028d"}],{}),e.cleanupOutdatedCaches(),e.registerRoute(new e.NavigationRoute(e.createHandlerBoundToURL("index.html"))),e.registerRoute(/(.*?)\.(js|css|ts)/,new e.CacheFirst({cacheName:"js-css-cache",plugins:[new e.ExpirationPlugin({maxEntries:100,maxAgeSeconds:604800}),new e.CacheableResponsePlugin({statuses:[0,200]})]}),"GET"),e.registerRoute(/(.*?)\.(png|jpe?g|svg|gif|bmp|psd|tiff|tga|eps|ico)/,new e.CacheFirst({cacheName:"image-cache",plugins:[new e.ExpirationPlugin({maxEntries:100,maxAgeSeconds:604800}),new e.CacheableResponsePlugin({statuses:[0,200]})]}),"GET")}));
