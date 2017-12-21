/*
  main.go
  
  version: 17.12.21
  Copyright (C) 2017 Jeroen P. Broks
  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.
  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:
  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/
package main
import (
	"github.com/jpbubble/base"
	"os"
	"path"
	"trickyunits/mkl"
	"trickyunits/qff"
	"trickyunits/jcr6/jcr6main"
_	"trickyunits/jcr6/jcr6realdir"
)

func init(){
mkl.Version("Bubble CLI - main.go","17.12.21")
mkl.Lic    ("Bubble CLI - main.go","ZLib License")
}

func main(){
	bubble.WriteLn("Amber","Bubble CLI\n")
	if len(os.Args)<3 {
		bubble.Write("Purple","Usage: ")
		bubble.Write("Yellow","bubble ")
		bubble.Write("Cyan","<JCR> <Entry>[.lua] ")
		bubble.WriteLn("LightBlue","[Parameters]")
		bubble.WriteLn("White","")
		bubble.WriteLn("Pink",mkl.ListAll())
		os.Exit(0)
	}
	jfile:=os.Args[1]
	jentry:=os.Args[2]
	if path.Ext(jentry)=="" { jentry+=".lua" }
	if !qff.Exists(jfile) { bubble.Fatal(jfile+": file not found!") }
	jcr:=bubble.SetJCR(jfile)
	if !jcr6main.HasEntry(jcr,jentry) { bubble.Fatal(jfile + ": "+jentry+": Entry not found") }
}
