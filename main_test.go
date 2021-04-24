package pornhubEvaler

import "testing"

var (
	dataOld = `var p=1992641534185; var s=389958454; var n;\nif ((s >> 13) & 1)      p+=/* 120886108*\n*/152880888*/* 120886108*\n*/14;/*\nelse p-=\n*/else\np-=/*\np+= */109108064*        14; if ((s >> 5) & 1)/*\np+= */p+=206490833*/*\nelse p-=\n*/6; else  p-=  303244837*/*\np+= */6;/*\nelse p-=\n*/if ((s >> 5) & 1) p+=/*\n*13;\n*/80475974*/* 120886108*\n*/6;    else    p-=179580231*/*\nelse p-=\n*/6;\nif ((s >> 4) & 1)/*\np+= */p+=/*\np+= */80455957* 7;/* 120886108*\n*/else p-=/* 120886108*\n*/25077917*/*\nelse p-=\n*/5; if ((s >> 15) & 1) p+=\n97957449*/*\np+= */16; else /*\n*13;\n*/p-=14479763*/*\n*13;\n*/16;\np-=2111472091;\nn=leastFactor(p);\n{ document.cookie="RNKEY="+n+"*"+p/n+":"+s+":2040500497:1";\n document.location.reload(true); }`
	dataNew = `var p=3050680868863; var s=359467023; var n;\nif ((s >> 13) & 1) p+=/*\n*13;\n*/63647793*/*\nelse p-=\n*/14;else /* 120886108*\n*/p-=\t35813149*/*\np+= */14;/*\nelse p-=\n*/if ((s >> 10) & 1)\np+=\n84837772*\n13;else p-= 10959757*\n11;\tif ((s >> 4) & 1) p+=\t240745649*\t7; else /* 120886108*\n*/p-= 74478472*\t5;/*\n*13;\n*/if ((s >> 8) & 1)\np+=\t83123550*/*\n*13;\n*/11;\telse p-=/*\nelse p-=\n*/202676694* 9;\tif ((s >> 8) & 1) p+= 94847454*\n11;else /*\n*13;\n*/p-=/*\nelse p-=\n*/98522317*\n9;/*\nelse p-=\n*/ p+=3161093800;\n n=leastFactor(p);\n{ document.cookie="RNKEY="+n+"*"+p/n+":"+s+":489684649:1; path=/";\n  document.location.reload(true); }`
)

func BenchmarkEvalWithOldData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Eval(dataOld)
	}
}

func BenchmarkEvalWithNewData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Eval(dataNew)
	}
}
