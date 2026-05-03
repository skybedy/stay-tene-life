document.querySelectorAll('[data-copy]').forEach(btn=>btn.addEventListener('click',()=>{const t=document.getElementById(btn.dataset.copy);navigator.clipboard.writeText(t.innerText)}));
document.querySelectorAll('[data-copy-text]').forEach(btn=>btn.addEventListener('click',()=>navigator.clipboard.writeText(btn.dataset.copyText)));
