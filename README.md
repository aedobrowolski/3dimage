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
be able to make the left and right columns overlap. Instead of two columns you
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
AOplfMR.bEAOplfMR.bEAOplfMR.bEAOplfMR.bEAOplfMR.bEAOplfMR.bEAOplfMR.bE
AGjf,pn+I.AGjf,pn+I.AGjf,pn+I.AGjf,pn+I.AGjf,pn+I.AGjf,pn+I.AGjf,pn+I.
AQOCb^.d#MAQOCb^.d#MAQOCb^.d#MAQOCb^.d#MAQOCb^.d#MAQOCb^.d#MAQOCb^.d#M
^.f+#lKR,Q^.f+#R,Q^.f+#R#R,Q^.f+#R#R,Q^ff+#R#R,Q^ff+#R#Q^ff+#R#Q#Q^ff+
Rh+@MCKIlARh+@MIlARh+@MIMIlARh+@MIMIlAh+@@MIMIlAh+@@MIMAh+@@MIMAMAh+@@
E^bpjMKdlIE^bpjdlIE^bpjdjdlIE^bpjdjdlE^bpjjdjdlE^bpjjdjE^bpjjdjEjE^bpj
A,pRfI#@d=A,pRf@d=A,pRf@f@d=A,pRf@f@=A,pRf@@f@=A,pRf@@fA,pRf@@fAfA,pRf
@fOh.A,QK+@fOh.QK+@fOh.Q.QK+@fOh.Q.K+@fOh.Q..K+@fOh.Q..@fOh.Q..@.@fOh.
h,bM@fOAl=h,bM@Al=h,bM@A@Al=h,bM@AAl=h,bM@AAll=h,bM@AAlh,bM@AAlhlh,bM@
h.fGbORjd,h.fGbjd,h.fGbjbjd,h.fGbbjd,h.fGbbjd,,h.fGbbjdh.fGbbjdhdh.fGb
Kfn=bhRMO+Kfn=bMO+Kfn=bMbMO+Kfn=MbMO+KKfnMbMO+KKKfnMbMOKKfnMbMOKOKKfnM
@EnlM!.hGR@EnlMhGR@EnlMhMhGR@EnMhMhGGR@EnMhhGGR@@EnMhhG@@EnMhhG@G@@EnM
#=b,.@AMEK#=b,.MEK#=b,.M.MEK#=,.M.MEEK#=,.MMEEK#==,.MME#==,.MME#E#==,.
p!QfjME@.dp!Qfj@.dp!Qfj@j@.dpQfj@j@@.dpQfj@j@.dpQffj@j@pQffj@j@p@pQffj
.G^@,fpMjE.G^@,MjE.G^@,M,MjE.G@,M,MjjE.G@,MMjjE.GG@,MMj.GG@,MMj.j.GG@,
@GO!d+=nfC@GO!dnfC@GO!dndnfC@GOdndnffC@GOdnnffC@@GOdnnf@@GOdnnf@f@@GOd
d#KCAGEQOnd#KCAQOnd#KCAQAQOnd#KCQAQOndd#KQAQOnddd#KQAQOdd#KQAQOdOdd#KQ
KdAn#lRh=IKdAn#h=IKdAn#h#h=IKdAn##h=IKdAn##h=IIKdAn##h=KdAn##h=K=KdAn#
pGlQKC,#IfpGlQK#IfpGlQK#K#IfpGlQK##IfpGlQK##IIfpGlQK##IpGlQK##IpIpGlQK
OnfI#l!+pMOnfI#+pMOnfI#+#+pMOnfI#+#pMOnfI#+##pMOnfI#+##OnfI#+##O#OnfI#
+fCjQGI^lR+fCjQ^lR+fCjQ^Q^lR+fCjQ^Q^R+fCjQ^^Q^R+fCjQ^^Q+fCjQ^^Q+Q+fCjQ
Af^l,hndIQAf^l,dIQAf^l,d,dIQAf^l,d,dIAf^l,,d,dIAf^l,,d,Af^l,,d,A,Af^l,
OK^,nCGh.dOK^,nh.dOK^,nhnh.dOK^,nhnh.dK^,,nhnh.dK^,,nhndK^,,nhndndK^,,
,jl^+@pIMh,jl^+IMh,jl^+I+IMh,jl^+I+IMh,ll^+I+IMh,ll^+I+h,ll^+I+h+h,ll^
!=^QM+.,Af!=^QM+.,Af!=^QM+.,Af!=^QM+.,Af!=^QM+.,Af!=^QM+.,Af!=^QM+.,Af
fC#QnOMb.=fC#QnOMb.=fC#QnOMb.=fC#QnOMb.=fC#QnOMb.=fC#QnOMb.=fC#QnOMb.=
CEbG#+^,=lCEbG#+^,=lCEbG#+^,=lCEbG#+^,=lCEbG#+^,=lCEbG#+^,=lCEbG#+^,=l
```

The input for the algorithm is a rectangle of digits.  The greater the 
digit the closer the plane will appear.  The background is represented
by the digit zero (or a blank).  Using Emacs picture mode is one way
of preparing these images.

```text
                                       .
  4       5       6                    .
     5              2    6      4      .
         3    6    212                 .
                  21112 3         3 5  .
   6     5       2111112      4        .
                211111112  6           .
 4     3     6 211  1  112             .
              2111  1  1112      5     .
        5    2 111  1 6111 2        3  .
   6           11111111111             .
            6  11111111111   3         .
 3             11122222111     4  5    .
       5       1112  62111  4          .
  6            1112   2111             .
           6   1112   2111             .
     5         5      6       4  3  5  .
```