# pornhubEvaler

### Explanation

When you made a lot of requests to pornhub.com, you've probably faced with the script, that protects website from further non browser requests

<details>
  <summary>The script</summary>
  
  ```javascript
  function leastFactor(n) {
if (isNaN(n) || !isFinite(n)) return NaN;
if (n==0) return 0;
if (n%1 || n*n<2) return 1;
if (n%2==0) return 2;
if (n%3==0) return 3;
if (n%5==0) return 5;
var m=Math.sqrt(n);
for (var i=7;i<=m;i+=30) {
 if (n%i==0)      return i;
 if (n%(i+4)==0)  return i+4;
 if (n%(i+6)==0)  return i+6;
 if (n%(i+10)==0) return i+10;
 if (n%(i+12)==0) return i+12;
 if (n%(i+16)==0) return i+16;
 if (n%(i+22)==0) return i+22;
 if (n%(i+24)==0) return i+24;
}
return n;
}
function go() {
var p=1992641534185; var s=389958454; var n;
if ((s >> 13) & 1)      p+=/* 120886108*
*/152880888*/* 120886108*
*/14;/*
else p-=
*/else
p-=/*
p+= */109108064*        14; if ((s >> 5) & 1)/*
p+= */p+=206490833*/*
else p-=
*/6; else  p-=  303244837*/*
p+= */6;/*
else p-=
*/if ((s >> 5) & 1) p+=/*
*13;
*/80475974*/* 120886108*
*/6;    else    p-=179580231*/*
else p-=
*/6;
if ((s >> 4) & 1)/*
p+= */p+=/*
p+= */80455957* 7;/* 120886108*
*/else p-=/* 120886108*
*/25077917*/*
else p-=
*/5; if ((s >> 15) & 1) p+=
97957449*/*
p+= */16; else /*
*13;
*/p-=14479763*/*
*13;
*/16;
p-=2111472091;
n=leastFactor(p);
{ document.cookie="RNKEY="+n+"*"+p/n+":"+s+":2040500497:1";
 document.location.reload(true); }
}
  ```
</details>

It called as a JS protection, and pornhub accept your requests only with correct RNKEY cookie
That lib helps you to generate this cookie even without js evaling
It uses regex to parse all needed numbers for further calculation

# Usage

```go
package main

import (
  "fmt"
  evaler "github.com/Numenorean/pornhubEvaler"
)

func main(){
  js := "long js code..." // You only need code inside "go" function
  cookie := evaler.Eval(js)
  fmt.Println(cookie)
}
```

# Benchmark

```
goos: windows
goarch: amd64
pkg: github.com/Numenorean/pornhubEvaler
cpu: Intel(R) Core(TM) i3-2120 CPU @ 3.30GHz
BenchmarkEvalWithOldData-4   	   26568	     40553 ns/op	    5324 B/op	      64 allocs/op
BenchmarkEvalWithNewData-4   	     297	   4054488 ns/op	    5498 B/op	      65 allocs/op
```
