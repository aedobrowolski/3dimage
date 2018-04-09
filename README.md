# 3D Images Done with Type #

We will generate three dimensional "magic eye" images with plain type.
The illusion will work if the person stares at the screen and tries
to line up equal letters.  Try to stare at the following two column
text and make the left and right columns coincide by focusing at a
distance.

```text
            It  is  a  period              It  is  a  period

            of civil war. Rebel          of civil war. Rebel

            spaceships   striking      spaceships   striking

            from a hidden base  ...  from a hidden base  ...
```

By staring at the blank space between columns and relaxing your eyes you should
be able to make the left and right columns overlap. Try zooming out in your
browser (e.g. 70%) if you cannot get overlap easily. Instead of two columns you
will see three columns and the middle column will appear in to contain text
that recedes in space.

The code in this project generates 3D images using the same principle. It will
repeat a pattern across a single rectangular field. Each line can start with a
new random pattern, but the base width of the pattern will always be the same.
For some runs in a line it will shorten the period of the pattern to create
closer objects. The smaller the period the closer the objects will appear. This
allows 3D images to pop out of the type when the eyes focus on some distant
spot in order to make adjacent cycles of the pattern overlap.

The text below shows one such generated image. See if you can see the hidden
objects.

```text
QAIjR+nfdK!lbO,#QAIjR+nfdK!lbO,#QAIjR+nfdK!lbO,#QAIjR+nfdK!lbO,#QAIjR+nfdK!l
EC!+M=hA.^@ObQlGEC!+M=hA.^@ObQlGEC!+M=hA.^@ObQlGEC!+M=hA.^@ObQlGEC!+M=hA.^@O
RGO^.EIKMp!d@b=,RGO^.EIKMp!d@b=,RGO^.EIKMp!d@b=,RGO^.EIKMp!d@b=,RGO^.EIKMp!d
Cpf+jA,^QMhI!KbdCpf+j^QMhI!KbdA,Cpf+j^QMhI!KbAd,Cpf+j^QMhI!Kb,Cpf+j^QMAdhI!K
I#OnCGQM+bEp^.A=I#OnCM+bEp^.A=GQI#OnCM+bEp^.=GQAI#OnCM+bEp^.=AI#OnCM+bGQEp^.
AIj^pd@.fnOh!=bEAIj^p.fnOh!=bEd@AIj^p.fnOh!bEd@A=Ij^p.fnOh!bEA=Ij^p.fnd@Oh!b
KI,=hMR.ljOdf+GCKI,=h.ljOdf+GCMRKI,=h.ljOd+GCMRKIf,=h.ljOd+GCKIf,=h.ljMROd+G
b.IEG=Qj+ndhR,f!b.IEGj+ndhR,f!=Qb.IEGj+ndR,f!=Qb.IhEGj+ndR,f!b.IhEGj+n=QdR,f
+G!@lMQO,#RA=npK+G!@lO,#RA=npKMQ+G!@lO,#A=npKMQ+G!@RlO,#A=npK+G!@RlO,#MQA=np
.^fCIbpKjQAlRG=h.^fCIKjQAlRG=hbp.^fCIKjAlRG=hbp.^fCIQKjAlRG=h.^fCIQKjAbplRG=
C@Oj=.E#Qf!,+dhlC@Oj=#Qf!,+dhl.EC@Oj=#f!,+dhQl.C@Oj=#Ef!,+dhQC@Oj=#Ef!l.,+dh
hG#bj!E.lpnKAIOQhG#bj.lpnKAIOQ!EhG#bjlpnKA.IOQ!Eh#bjlpGnKA.IOEh#bjlpGnQ!KA.I
h#IRMfEjK@+GnCA^h#IRMjK@+GnCA^fEh#IRjK@+GnMCA^fEhIRjK@+#GnMCAEhIRjK@+#^fGnMC
+lQRbh^G!AMECd=K+lQRbG!AMECd=Kh^+lQbG!AMERCd=Kh^+lbG!AMEQRCd=^+lbG!AMEKhQRCd
dhRAG+M^fK@blIE=dhRAG^fK@blIE=+MdhRA^fK@blGIE=+MdRA^fK@hblGIEMdRA^fK@h=+blGI
ME!#ROGbnj.lQC^AME!#Rbnj.lQC^AOGME!#Rnj.lQbC^AOGM!#RnjE.lQbC^GM!#RnjE.AOlQbC
.h+,ACOjpQdlMn#!.h+,AjpQdlMn#!CO.h+,AjQdlMn#p!C.h+,AjOQdlMn#p.h+,AjOQd!ClMn#
#K^=G,b+IEnQjlRf#K^=G+IEnQjlRf,b#K^=G+InQjlRf,b#K^=GE+InQjlRf#K^=GE+In,bQjlR
!fbnjd+MI,GQ.#K@!fbnjMI,GQ.#K@d+!fbnjMI,Q.#K@d+!fbnGjMI,Q.#K@!fbnGjMI,d+Q.#K
!#R@.OlpfMdbnAEh!#R@.pfMdbnAEhOl!#R@.pfMdnAEhOl!#Rb@.pfMdnAEh!#Rb@.pfMOldnAE
MGjnl+hdAE@^O=Q.MGjnldAE@^O=Q.+hMGjnldAE@^=Q.+hMGOjnldAE@^=Q.MGOjnldAE+h@^=Q
InM=Gf^lAK@hdO#EInM=GlAK@hdO#Ef^InM=GlAK@hd#Ef^IOnM=GlAK@hd#EIOnM=GlAKf^@hd#
!+,Kn^QGlOjRCf=p!+,KnGlOjRCf=p^Q!+,KnGlOjRCfp^Q=!+,KnGlOjRCfp=!+,KnGlO^QjRCf
KjlQId=fE.hbCMp^KjlQIfE.hbCMp^d=KjlQIfE.hbCMpd^=KjlQIfE.hbCMp=KjlQIfE.d^hbCM
,MIpRCnEQ+.KGdfj,MIpRCnEQ+.KGdfj,MIpRCnEQ+.KGdfj,MIpRCnEQ+.KGdfj,MIpRCnEQ+.K
fAb!Odj^GI.@M=plfAb!Odj^GI.@M=plfAb!Odj^GI.@M=plfAb!Odj^GI.@M=plfAb!Odj^GI.@
f@Qd=,GIlK#+jph.f@Qd=,GIlK#+jph.f@Qd=,GIlK#+jph.f@Qd=,GIlK#+jph.f@Qd=,GIlK#+
```

The input for the algorithm is a rectangle of digits. The greater the digit the
closer the plane will appear in 3D. The background is represented by the digit
zero (or a blank). Using Emacs picture mode is one way of preparing these
images. Below is the pattern for a house.

```text
                                        
                                        
                    2                   
                   222  111             
                  22222 111             
         1       2222222111             
        111     22222222211             
       11111   222     2221             
      1111111 2222     2222             
     11     122222     22222            
    111     1222222222222222            
    111     12222222222222221111111     
    11111111122222     222221111111     
    11111111122222     222221111111     
    11111111122222     222221111111     
    11111111122222     222221111111     
33333333333333333333333333333333333333333
```

The output, using a 16 character wide random pattern is this:

```text
^bKhE,p.=nIOGQfj^bKhE,p.=nIOGQfj^bKhE,p.=nIOGQfj^bKhE,p.
=db#M.EnA@GpjOhQ=db#M.EnA@GpjOhQ=db#M.EnA@GpjOhQ=db#M.En
G^h+EO!nd@RKjlMbG^h+EO!nd@RKjlMbG^h+!OEnd@RKjlMbG^h+!OEn
M#GhInKCp+dA@,E.M#GhInKCp+dA@,E.M#GnKChI+dAp@,E.M#GnKChI
=f^#AM+,KhdjGp.I=f^#AM+,KhdjGp.I=fAM+,K^hdj#Gp.I=fAM+,K^
,O=Q#+GKlI.hEMf^,O=Q#+GKl.IhEMf^,Q#+GKl.OIh=EMf^,Q#+GKl.
fl!M.jQCKRG@EbO,fl!M.jQCRG@KEbO,!M.jQCRG@fKlEbO,!M.jQCRG
#hKGM@dAI,+njR.E#hKGM@dI,+njAR.hKGE#M@d+njI,AR.hKGE#M@d+
RC+IOjKlh^Q#!EG.RC+IOjlh^Q#!EKRC+IG.OjlQ#!Eh^KRC+IG.OjlQ
MjbdKOA,fp=#^l+IMjbdKA,Ofp=#lIMjbd^+KA,p=#lIOfMjbd^+KA,p
b,+CfMR!Gd@phIE=b,+CMR!fGd@pI=b,+CMR!fGd@pI=hEb,+CMR!fGd
!d@bO.#GKnQCEIp+!d@b.#GOKnQCI+!d@b.#GOKnQCI+E!d@b.#pGOKn
pK^C=+O#QlEn.GdbpK^C+O#QlEn.GbpK^C=d+O#En.GbQpK^C=dl+O#E
p^h.Gj#,QM=l!CnOp^h.j#,QM=l!COp^h.Gnj#,=l!COQp^h.GnMj#,=
GdQbI=OfjE.C,A@KGdQb=OfjE.C,AKGdQbI@=Of.C,AKjGdQbI@E=Of.
fMQA#=RC.I!E,K^bfMQA=RC.I!E,KbfMQA#^=RC!E,Kb.fMQA#^I=RC!
b=ldC#fAn.K@Ej^,dC#fAn.K@Ej^,dC#fAn.K@Ej^,dC#fAn.K@Ej^,d
```

To build the exe you must have go installed on your system. Then cd into the
cmd/3dimage folder and type "go build". You will find a new binary/exe called
3dimage in the folder. Running the program with no arguments shows the help.

I welcome pull requests for your own samples which I can put into the samples
folder. Hope you have as much fun playing with this project as my 12 year old
son and I had building it.