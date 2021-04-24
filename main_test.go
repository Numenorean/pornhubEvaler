package pornhubEvaler

import "testing"

var data = `var p=1992641534185; var s=389958454; var n;
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
 document.location.reload(true); }`

func BenchmarkEval(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Eval(data)
	}
}
