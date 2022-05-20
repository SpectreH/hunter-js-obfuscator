package hunterjsobfuscator

import (
	"math/rand"
	"testing"
)

func TestObfuscate(t *testing.T) {
	rand.Seed(1000)

	expected := `var _0xc3e=["","split","0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/","slice","indexOf","","",".","pow","reduce","reverse","0"];function _0xe82c(d,e,f){var g=_0xc3e[2][_0xc3e[1]](_0xc3e[0]);var h=g[_0xc3e[3]](0,e);var i=g[_0xc3e[3]](0,f);var j=d[_0xc3e[1]](_0xc3e[0])[_0xc3e[10]]()[_0xc3e[9]](function(a,b,c){if(h[_0xc3e[4]](b)!==-1)return a+=h[_0xc3e[4]](b)*(Math[_0xc3e[8]](e,c))},0);var k=_0xc3e[0];while(j>0){k=i[j%f]+k;j=(j-(j%f))/f}return k||_0xc3e[11]}eval(function(h,u,n,t,e,r){r="";for(var i=0,len=h.length;i<len;i++){var s="";while(h[i]!==n[e]){s+=h[i];i++}for(var j=0;j<n.length;j++)s=s.replace(new RegExp(n[j],"g"),j);r+=String.fromCharCode(_0xe82c(s,e,10)-t)}return decodeURIComponent(escape(r))}("ffXXONfffffNffXOXNffOXfNffOfXNfOOONfOOfNfffXXNffXOXNfffffNfffffNfffOfNfOOfNOXXXNOOXXN",64,"XfONStpRI",13,3,33))`

	o := NewObfuscator("alert('hello');")
	code := o.Obfuscate()

	if expected != code {
		t.Errorf("Objuscated code is not the same as expected")
	}
}
