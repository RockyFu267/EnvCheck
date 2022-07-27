package gpuid

var GPUMap string = `
0008	NV1 [STG2000X-B Series]	
0009	NV1 [NV1 Series]	
000b		
0010		
0018	NV3 [Riva 128]	
0019	NV3 [Riva 128ZX]	
0020	NV4 [Riva TNT]	
0028	NV5 [Riva TNT2 / TNT2 Pro]	
0029	NV5 [Riva TNT2 Ultra]	
002a	NV5 [Riva TNT2]	
002b	NV5 [Riva TNT2]	
002c	NV5 [Vanta / Vanta LT]	
002d	NV5 [Riva TNT2 Model 64 / Model 64 Pro]	
002e		
002f		
0034	MCP04 SMBus	
0035	MCP04 IDE	
0036	MCP04 Serial ATA Controller	
0037	MCP04 Ethernet Controller	
0038	MCP04 Ethernet Controller	
003a	MCP04 AC'97 Audio Controller	
003b	MCP04 USB Controller	
003c	MCP04 USB Controller	
003d	MCP04 PCI Bridge	
003e	MCP04 Serial ATA Controller	
0040	NV40 [GeForce 6800 Ultra]	
0041	NV40 [GeForce 6800]	
0042	NV40 [GeForce 6800 LE]	
0043	NV40 [GeForce 6800 XE]	
0044	NV40 [GeForce 6800 XT]	
0045	NV40 [GeForce 6800 GT]	
0046		
0047	NV40 [GeForce 6800 GS]	
0048	NV40 [GeForce 6800 XT]	
0049		
004d		
004e	NV40GL [Quadro FX 4000]	
0050	CK804 ISA Bridge	
0051	CK804 ISA Bridge	
0052	CK804 SMBus	
0053	CK804 IDE	
0054	CK804 Serial ATA Controller	
0055	CK804 Serial ATA Controller	
0056	CK804 Ethernet Controller	
0057	CK804 Ethernet Controller	
0058	CK804 AC'97 Modem	
0059	CK804 AC'97 Audio Controller	
005a	CK804 USB Controller	
005b	CK804 USB Controller	
005c	CK804 PCI Bridge	
005d	CK804 PCIE Bridge	
005e	CK804 Memory Controller	
005f	CK804 Memory Controller	
0060	nForce2 ISA Bridge	
0064	nForce2 SMBus (MCP)	
0065	nForce2 IDE	
0066	nForce2 Ethernet Controller	
0067	nForce2 USB Controller	
0068	nForce2 USB Controller	
006a	nForce2 AC97 Audio Controler (MCP)	
006b	nForce Audio Processing Unit	
006c	nForce2 External PCI Bridge	
006d	nForce2 PCI Bridge	
006e	nForce2 FireWire (IEEE 1394) Controller	
0080	MCP2A ISA bridge	
0084	MCP2A SMBus	
0085	MCP2A IDE	
0086	MCP2A Ethernet Controller	
0087	MCP2A USB Controller	
0088	MCP2A USB Controller	
008a	MCP2S AC'97 Audio Controller	
008b	MCP2A PCI Bridge	
008c	MCP2A Ethernet Controller	
008e	nForce2 Serial ATA Controller	
0090	G70 [GeForce 7800 GTX]	
0091	G70 [GeForce 7800 GTX]	
0092	G70 [GeForce 7800 GT]	
0093	G70 [GeForce 7800 GS]	
0095	G70 [GeForce 7800 SLI]	
0097	G70 [GeForce GTS 250]	
0098	G70M [GeForce Go 7800]	
0099	G70M [GeForce Go 7800 GTX]	
009d	G70GL [Quadro FX 4500]	
00a0	NV0A [Aladdin TNT2 IGP]	
00c0	NV41 [GeForce 6800 GS]	
00c1	NV41 [GeForce 6800]	
00c2	NV41 [GeForce 6800 LE]	
00c3	NV41 [GeForce 6800 XT]	
00c5	NV41	
00c6	NV41	
00c7	NV41	
00c8	NV41M [GeForce Go 6800]	
00c9	NV41M [GeForce Go 6800 Ultra]	
00cc	NV41GLM [Quadro FX Go1400]	
00cd	NV42GL [Quadro FX 3450/4000 SDI]	
00ce	NV41GL [Quadro FX 1400]	
00cf	NV41	
00d0	nForce3 LPC Bridge	
00d1	nForce3 Host Bridge	
00d2	nForce3 AGP Bridge	
00d3	CK804 Memory Controller	
00d4	nForce3 SMBus	
00d5	nForce3 IDE	
00d6	nForce3 Ethernet	
00d7	nForce3 USB 1.1	
00d8	nForce3 USB 2.0	
00d9	nForce3 Audio	
00da	nForce3 Audio	
00dd	nForce3 PCI Bridge	
00df	CK8S Ethernet Controller	
00e0	nForce3 250Gb LPC Bridge	
00e1	nForce3 250Gb Host Bridge	
00e2	nForce3 250Gb AGP Host to PCI Bridge	
00e3	nForce3 Serial ATA Controller	
00e4	nForce 250Gb PCI System Management	
00e5	CK8S Parallel ATA Controller (v2.5)	
00e6	CK8S Ethernet Controller	
00e7	CK8S USB Controller	
00e8	nForce3 EHCI USB 2.0 Controller	
00ea	nForce3 250Gb AC'97 Audio Controller	
00ed	nForce3 250Gb PCI-to-PCI Bridge	
00ee	nForce3 Serial ATA Controller 2	
00f0		
00f1	NV43 [GeForce 6600 GT]	
00f2	NV43 [GeForce 6600]	
00f3	NV43 [GeForce 6200]	
00f4	NV43 [GeForce 6600 LE]	
00f5	G70/G71 [GeForce 7800 GS AGP]	
00f6	NV43 [GeForce 6800 GS/XT]	
00f8	NV45GL [Quadro FX 3400/4400]	
00f9	NV40 [GeForce 6800 GT/GTO/Ultra]	
00fa	NV39 [GeForce PCX 5750]	
00fb	NV35 [GeForce PCX 5900]	
00fc	NV37GL [Quadro FX 330/GeForce PCX 5300]	
00fd	NV37GL [Quadro PCI-E Series]	
00fe	NV38GL [Quadro FX 1300]	
00ff	NV19 [GeForce PCX 4300]	
0100	NV10 [GeForce 256 SDR]	
0101	NV10 [GeForce 256 DDR]	
0103	NV10GL [Quadro]	
0110	NV11 [GeForce2 MX/MX 400]	
0111	NV11 [GeForce2 MX200]	
0112	NV11M [GeForce2 Go]	
0113	NV11GL [Quadro2 MXR/EX/Go]	
0140	NV43 [GeForce 6600 GT]	
0141	NV43 [GeForce 6600]	
0142	NV43 [GeForce 6600 LE]	
0143	NV43 [GeForce 6600 VE]	
0144	NV43M [GeForce Go 6600]	
0145	NV43 [GeForce 6610 XL]	
0146	NV43M [GeForce Go6200 TE / 6600 TE]	
0147	NV43 [GeForce 6700 XL]	
0148	NV43M [GeForce Go 6600]	
0149	NV43M [GeForce Go 6600 GT]	
014a	NV43 [Quadro NVS 440]	
014b	NV43	
014c		
014d	NV43GL [Quadro FX 550]	
014e	NV43GL [Quadro FX 540]	
014f	NV43 [GeForce 6200]	
0150	NV15 [GeForce2 GTS/Pro]	
0151	NV15 [GeForce2 Ti]	
0152	NV15 [GeForce2 Ultra]	
0153	NV15GL [Quadro2 Pro]	
0160	NV44 [GeForce 6500]	
0161	NV44 [GeForce 6200 TurboCache]	
0162	NV44 [GeForce 6200 SE TurboCache]	
0163	NV44 [GeForce 6200 LE]	
0164	NV44M [GeForce Go 6200]	
0165	NV44 [Quadro NVS 285]	
0166	NV44M [GeForce Go 6400]	
0167	NV44M [GeForce Go 6200]	
0168	NV44M [GeForce Go 6400]	
0169	NV44 [GeForce 6250]	
016a	NV44 [GeForce 7100 GS]	
016d	NV44	
016e	NV44	
016f	NV44	
0170	NV17 [GeForce4 MX 460]	
0171	NV17 [GeForce4 MX 440]	
0172	NV17 [GeForce4 MX 420]	
0173	NV17 [GeForce4 MX 440-SE]	
0174	NV17M [GeForce4 440 Go]	
0175	NV17M [GeForce4 420 Go]	
0176	NV17M [GeForce4 420 Go 32M]	
0177	NV17M [GeForce4 460 Go]	
0178	NV17GL [Quadro4 550 XGL]	
0179	NV17M [GeForce4 440 Go 64M]	
017a	NV17GL [Quadro NVS]	
017b	NV17GL [Quadro4 550 XGL]	
017c	NV17GL [Quadro4 500 GoGL]	
017d		
017f	NV17	
0181	NV18 [GeForce4 MX 440 AGP 8x]	
0182	NV18 [GeForce4 MX 440SE AGP 8x]	
0183	NV18 [GeForce4 MX 420 AGP 8x]	
0184	NV18 [GeForce4 MX]	
0185	NV18 [GeForce4 MX 4000]	
0186	NV18M [GeForce4 448 Go]	
0187	NV18M [GeForce4 488 Go]	
0188	NV18GL [Quadro4 580 XGL]	
0189	NV18 [GeForce4 MX with AGP8X (Mac)]	
018a	NV18GL [Quadro NVS 280 SD]	
018b	NV18GL [Quadro4 380 XGL]	
018c	NV18GL [Quadro NVS 50 PCI]	
018d	NV18M [GeForce4 448 Go]	
018f		
0190	G80 [GeForce 8800 GTS / 8800 GTX]	
0191	G80 [GeForce 8800 GTX]	
0192	G80 [GeForce 8800 GTS]	
0193	G80 [GeForce 8800 GTS]	
0194	G80 [GeForce 8800 Ultra]	
0197	G80GL [Tesla C870]	
019d	G80GL [Quadro FX 5600]	
019e	G80GL [Quadro FX 4600]	
01a0	nForce 220/420 NV1A [GeForce2 MX]	
01a4	nForce CPU bridge	
01ab	nForce 420 Memory Controller (DDR)	
01ac	nForce 220/420 Memory Controller	
01ad	nForce 220/420 Memory Controller	
01b0	nForce Audio Processing Unit	
01b1	nForce AC'97 Audio Controller	
01b2	nForce ISA Bridge	
01b4	nForce PCI System Management	
01b7	nForce AGP to PCI Bridge	
01b8	nForce PCI-to-PCI bridge	
01bc	nForce IDE	
01c1	nForce AC'97 Modem Controller	
01c2	nForce USB Controller	
01c3	nForce Ethernet Controller	
01d0	G72 [GeForce 7350 LE]	
01d1	G72 [GeForce 7300 LE]	
01d2	G72 [GeForce 7550 LE]	
01d3	G72 [GeForce 7200 GS / 7300 SE]	
01d5	G72	
01d6	G72M [GeForce Go 7200]	
01d7	G72M [Quadro NVS 110M/GeForce Go 7300]	
01d8	G72M [GeForce Go 7400]	
01d9	G72M [GeForce Go 7450]	
01da	G72M [Quadro NVS 110M]	
01db	G72M [Quadro NVS 120M]	
01dc	G72GLM [Quadro FX 350M]	
01dd	G72 [GeForce 7500 LE]	
01de	G72GL [Quadro FX 350]	
01df	G72 [GeForce 7300 GS]	
01e0	nForce2 IGP2	
01e8	nForce2 AGP	
01ea	nForce2 Memory Controller 0	
01eb	nForce2 Memory Controller 1	
01ec	nForce2 Memory Controller 2	
01ed	nForce2 Memory Controller 3	
01ee	nForce2 Memory Controller 4	
01ef	nForce2 Memory Controller 5	
01f0	NV1F C17 [GeForce4 MX IGP]	
0200	NV20 [GeForce3]	
0201	NV20 [GeForce3 Ti 200]	
0202	NV20 [GeForce3 Ti 500]	
0203	NV20GL [Quadro DCC]	
0211	NV48 [GeForce 6800]	
0212	NV48 [GeForce 6800 LE]	
0215	NV48 [GeForce 6800 GT]	
0218	NV48 [GeForce 6800 XT]	
0221	NV44A [GeForce 6200]	
0222	NV44 [GeForce 6200 A-LE]	
0224	NV44	
0240	C51PV [GeForce 6150]	
0241	C51 [GeForce 6150 LE]	
0242	C51G [GeForce 6100]	
0243	C51 PCI Express Bridge	
0244	C51 [GeForce Go 6150]	
0245	C51 [Quadro NVS 210S/GeForce 6150LE]	
0246	C51 PCI Express Bridge	
0247	C51 [GeForce Go 6100]	
0248	C51 PCI Express Bridge	
0249	C51 PCI Express Bridge	
024a	C51 PCI Express Bridge	
024b	C51 PCI Express Bridge	
024c	C51 PCI Express Bridge	
024d	C51 PCI Express Bridge	
024e	C51 PCI Express Bridge	
024f	C51 PCI Express Bridge	
0250	NV25 [GeForce4 Ti 4600]	
0251	NV25 [GeForce4 Ti 4400]	
0252	NV25 [GeForce4 Ti]	
0253	NV25 [GeForce4 Ti 4200]	
0258	NV25GL [Quadro4 900 XGL]	
0259	NV25GL [Quadro4 750 XGL]	
025b	NV25GL [Quadro4 700 XGL]	
0260	MCP51 LPC Bridge	
0261	MCP51 LPC Bridge	
0262	MCP51 LPC Bridge	
0263	MCP51 LPC Bridge	
0264	MCP51 SMBus	
0265	MCP51 IDE	
0266	MCP51 Serial ATA Controller	
0267	MCP51 Serial ATA Controller	
0268	MCP51 Ethernet Controller	
0269	MCP51 Ethernet Controller	
026a	MCP51 MCI	
026b	MCP51 AC97 Audio Controller	
026c	MCP51 High Definition Audio	
026d	MCP51 USB Controller	
026e	MCP51 USB Controller	
026f	MCP51 PCI Bridge	
0270	MCP51 Host Bridge	
0271	MCP51 PMU	
0272	MCP51 Memory Controller 0	
027e	C51 Memory Controller 2	
027f	C51 Memory Controller 3	
0280	NV28 [GeForce4 Ti 4800]	
0281	NV28 [GeForce4 Ti 4200 AGP 8x]	
0282	NV28 [GeForce4 Ti 4800 SE]	
0286	NV28M [GeForce4 Ti 4200 Go AGP 8x]	
0288	NV28GL [Quadro4 980 XGL]	
0289	NV28GL [Quadro4 780 XGL]	
028c	NV28GLM [Quadro4 Go700]	
0290	G71 [GeForce 7900 GTX]	
0291	G71 [GeForce 7900 GT/GTO]	
0292	G71 [GeForce 7900 GS]	
0293	G71 [GeForce 7900 GX2]	
0294	G71 [GeForce 7950 GX2]	
0295	G71 [GeForce 7950 GT]	
0297	G71M [GeForce Go 7950 GTX]	
0298	G71M [GeForce Go 7900 GS]	
0299	G71M [GeForce Go 7900 GTX]	
029a	G71GLM [Quadro FX 2500M]	
029b	G71GLM [Quadro FX 1500M]	
029c	G71GL [Quadro FX 5500]	
029d	G71GL [Quadro FX 3500]	
029e	G71GL [Quadro FX 1500]	
029f	G71GL [Quadro FX 4500 X2]	
02a0	NV2A [XGPU]	Xbox Graphics Processing Unit (Integrated). GeForce3 derivative (NV20 < NV2A < NV25).
02a5	MCPX CPU Bridge	
02a6	MCPX Memory Controller	
02e0	G73 [GeForce 7600 GT AGP]	
02e1	G73 [GeForce 7600 GS AGP]	
02e2	G73 [GeForce 7300 GT AGP]	
02e3	G71 [GeForce 7900 GS AGP]	
02e4	G71 [GeForce 7950 GT AGP]	
02e5	G71 [GeForce 7600 GS AGP]	
02f0	C51 Host Bridge	
02f1	C51 Host Bridge	
02f2	C51 Host Bridge	
02f3	C51 Host Bridge	
02f4	C51 Host Bridge	
02f5	C51 Host Bridge	
02f6	C51 Host Bridge	
02f7	C51 Host Bridge	
02f8	C51 Memory Controller 5	
02f9	C51 Memory Controller 4	
02fa	C51 Memory Controller 0	
02fb	C51 PCI Express Bridge	
02fc	C51 PCI Express Bridge	
02fd	C51 PCI Express Bridge	
02fe	C51 Memory Controller 1	
02ff	C51 Host Bridge	
0300	NV30 [GeForce FX]	
0301	NV30 [GeForce FX 5800 Ultra]	
0302	NV30 [GeForce FX 5800]	
0308	NV30GL [Quadro FX 2000]	
0309	NV30GL [Quadro FX 1000]	
0311	NV31 [GeForce FX 5600 Ultra]	
0312	NV31 [GeForce FX 5600]	
0313		
0314	NV31 [GeForce FX 5600XT]	
0316	NV31M	
0317		
0318	NV31GL	
031a	NV31M [GeForce FX Go5600]	
031b	NV31M [GeForce FX Go5650]	
031c	NV31GLM [Quadro FX Go700]	
031d		
031e		
031f		
0320	NV34 [GeForce FX 5200]	
0321	NV34 [GeForce FX 5200 Ultra]	
0322	NV34 [GeForce FX 5200]	
0323	NV34 [GeForce FX 5200LE]	
0324	NV34M [GeForce FX Go5200 64M]	
0325	NV34M [GeForce FX Go5250]	
0326	NV34 [GeForce FX 5500]	
0327	NV34 [GeForce FX 5100]	
0328	NV34M [GeForce FX Go5200 32M/64M]	
0329	NV34M [GeForce FX Go5200]	
032a	NV34GL [Quadro NVS 280 PCI]	
032b	NV34GL [Quadro FX 500/600 PCI]	
032c	NV34M [GeForce FX Go5300 / Go5350]	
032d	NV34M [GeForce FX Go5100]	
032e	NV34	
032f	NV34 [GeForce FX 5200]	
0330	NV35 [GeForce FX 5900 Ultra]	
0331	NV35 [GeForce FX 5900]	
0332	NV35 [GeForce FX 5900XT]	
0333	NV38 [GeForce FX 5950 Ultra]	
0334	NV35 [GeForce FX 5900ZT]	
0338	NV35GL [Quadro FX 3000]	
033f	NV35GL [Quadro FX 700]	
0341	NV36 [GeForce FX 5700 Ultra]	
0342	NV36 [GeForce FX 5700]	
0343	NV36 [GeForce FX 5700LE]	
0344	NV36 [GeForce FX 5700VE]	
0345		
0347	NV36M [GeForce FX Go5700]	
0348	NV36M [GeForce FX Go5700]	
0349		
034b		
034c	NV36 [Quadro FX Go1000]	
034d	NV36	
034e	NV36GL [Quadro FX 1100]	
034f		
0360	MCP55 LPC Bridge	
0361	MCP55 LPC Bridge	
0362	MCP55 LPC Bridge	
0363	MCP55 LPC Bridge	
0364	MCP55 LPC Bridge	
0365	MCP55 LPC Bridge	
0366	MCP55 LPC Bridge	
0367	MCP55 LPC Bridge	
0368	MCP55 SMBus Controller	
0369	MCP55 Memory Controller	
036a	MCP55 Memory Controller	
036b	MCP55 SMU	
036c	MCP55 USB Controller	
036d	MCP55 USB Controller	
036e	MCP55 IDE	
036f		
0370	MCP55 PCI bridge	
0371	MCP55 High Definition Audio	
0372	MCP55 Ethernet	
0373	MCP55 Ethernet	
0374	MCP55 PCI Express bridge	
0375	MCP55 PCI Express bridge	
0376	MCP55 PCI Express bridge	
0377	MCP55 PCI Express bridge	
0378	MCP55 PCI Express bridge	
037a	MCP55 Memory Controller	
037c		
037e	MCP55 SATA Controller	
037f	MCP55 SATA Controller	
038b	G73 [GeForce 7650 GS]	
0390	G73 [GeForce 7650 GS]	
0391	G73 [GeForce 7600 GT]	
0392	G73 [GeForce 7600 GS]	
0393	G73 [GeForce 7300 GT]	
0394	G73 [GeForce 7600 LE]	
0395	G73 [GeForce 7300 GT]	
0396	G73	
0397	G73M [GeForce Go 7700]	
0398	G73M [GeForce Go 7600]	
0399	G73M [GeForce Go 7600 GT]	
039a	G73M [Quadro NVS 300M]	
039b	G73M [GeForce Go 7900 SE]	
039c	G73GLM [Quadro FX 550M]	
039d	G73	
039e	G73GL [Quadro FX 560]	
039f	G73	
03a0	C55 Host Bridge	
03a1	C55 Host Bridge	
03a2	C55 Host Bridge	
03a3	C55 Host Bridge	
03a4	C55 Host Bridge	
03a5	C55 Host Bridge	
03a6	C55 Host Bridge	
03a7	C55 Host Bridge	
03a8	C55 Memory Controller	
03a9	C55 Memory Controller	
03aa	C55 Memory Controller	
03ab	C55 Memory Controller	
03ac	C55 Memory Controller	
03ad	C55 Memory Controller	
03ae	C55 Memory Controller	
03af	C55 Memory Controller	
03b0	C55 Memory Controller	
03b1	C55 Memory Controller	
03b2	C55 Memory Controller	
03b3	C55 Memory Controller	
03b4	C55 Memory Controller	
03b5	C55 Memory Controller	
03b6	C55 Memory Controller	
03b7	C55 PCI Express bridge	
03b8	C55 PCI Express bridge	
03b9	C55 PCI Express bridge	
03ba	C55 Memory Controller	
03bb	C55 PCI Express bridge	
03bc	C55 Memory Controller	
03d0	C61 [GeForce 6150SE nForce 430]	
03d1	C61 [GeForce 6100 nForce 405]	
03d2	C61 [GeForce 6100 nForce 400]	
03d5	C61 [GeForce 6100 nForce 420]	
03d6	C61 [GeForce 7025 / nForce 630a]	
03e0	MCP61 LPC Bridge	
03e1	MCP61 LPC Bridge	
03e2	MCP61 Host Bridge	
03e3	MCP61 LPC Bridge	
03e4	MCP61 High Definition Audio	
03e5	MCP61 Ethernet	
03e6	MCP61 Ethernet	
03e7	MCP61 SATA Controller	
03e8	MCP61 PCI Express bridge	
03e9	MCP61 PCI Express bridge	
03ea	MCP61 Memory Controller	
03eb	MCP61 SMBus	
03ec	MCP61 IDE	
03ee	MCP61 Ethernet	
03ef	MCP61 Ethernet	
03f0	MCP61 High Definition Audio	
03f1	MCP61 USB 1.1 Controller	
03f2	MCP61 USB 2.0 Controller	
03f3	MCP61 PCI bridge	
03f4	MCP61 SMU	
03f5	MCP61 Memory Controller	
03f6	MCP61 SATA Controller	
03f7	MCP61 SATA Controller	
0400	G84 [GeForce 8600 GTS]	
0401	G84 [GeForce 8600 GT]	
0402	G84 [GeForce 8600 GT]	
0403	G84 [GeForce 8600 GS]	
0404	G84 [GeForce 8400 GS]	
0405	G84M [GeForce 9500M GS]	
0406	G84 [GeForce 8300 GS]	
0407	G84M [GeForce 8600M GT]	
0408	G84M [GeForce 9650M GS]	
0409	G84M [GeForce 8700M GT]	
040a	G84GL [Quadro FX 370]	
040b	G84GLM [Quadro NVS 320M]	
040c	G84GLM [Quadro FX 570M]	
040d	G84GLM [Quadro FX 1600M]	
040e	G84GL [Quadro FX 570]	
040f	G84GL [Quadro FX 1700]	
0410	G92 [GeForce GT 330]	
0414	G92 [GeForce 9800 GT]	
0418	G92 [GeForce GT 330 OEM]	
0420	G86 [GeForce 8400 SE]	
0421	G86 [GeForce 8500 GT]	
0422	G86 [GeForce 8400 GS]	
0423	G86 [GeForce 8300 GS]	
0424	G86 [GeForce 8400 GS]	
0425	G86M [GeForce 8600M GS]	
0426	G86M [GeForce 8400M GT]	
0427	G86M [GeForce 8400M GS]	
0428	G86M [GeForce 8400M G]	
0429	G86M [Quadro NVS 140M]	
042a	G86M [Quadro NVS 130M]	
042b	G86M [Quadro NVS 135M]	
042c	G86 [GeForce 9400 GT]	
042d	G86GLM [Quadro FX 360M]	
042e	G86M [GeForce 9300M G]	
042f	G86 [Quadro NVS 290]	
0440	MCP65 LPC Bridge	
0441	MCP65 LPC Bridge	
0442	MCP65 LPC Bridge	
0443	MCP65 LPC Bridge	
0444	MCP65 Memory Controller	
0445	MCP65 Memory Controller	
0446	MCP65 SMBus	
0447	MCP65 SMU	
0448	MCP65 IDE	
0449	MCP65 PCI bridge	
044a	MCP65 High Definition Audio	
044b	MCP65 High Definition Audio	
044c	MCP65 AHCI Controller	
044d	MCP65 AHCI Controller	
044e	MCP65 AHCI Controller	
044f	MCP65 AHCI Controller	
0450	MCP65 Ethernet	
0451	MCP65 Ethernet	
0452	MCP65 Ethernet	
0453	MCP65 Ethernet	
0454	MCP65 USB 1.1 OHCI Controller	
0455	MCP65 USB 2.0 EHCI Controller	
0456	MCP65 USB Controller	
0457	MCP65 USB Controller	
0458	MCP65 PCI Express bridge	
0459	MCP65 PCI Express bridge	
045a	MCP65 PCI Express bridge	
045b	MCP65 PCI Express bridge	
045c	MCP65 SATA Controller	
045d	MCP65 SATA Controller	
045e	MCP65 SATA Controller	
045f	MCP65 SATA Controller	
0531	C67 [GeForce 7150M / nForce 630M]	
0533	C67 [GeForce 7000M / nForce 610M]	
053a	C68 [GeForce 7050 PV / nForce 630a]	
053b	C68 [GeForce 7050 PV / nForce 630a]	
053e	C68 [GeForce 7025 / nForce 630a]	
0541	MCP67 Memory Controller	
0542	MCP67 SMBus	
0543	MCP67 Co-processor	
0547	MCP67 Memory Controller	
0548	MCP67 ISA Bridge	
054c	MCP67 Ethernet	
054d	MCP67 Ethernet	
054e	MCP67 Ethernet	
054f	MCP67 Ethernet	
0550	MCP67 AHCI Controller	
0554	MCP67 AHCI Controller	
0555	MCP67 SATA Controller	
055c	MCP67 High Definition Audio	
055d	MCP67 High Definition Audio	
055e	MCP67 OHCI USB 1.1 Controller	
055f	MCP67 EHCI USB 2.0 Controller	
0560	MCP67 IDE Controller	
0561	MCP67 PCI Bridge	
0562	MCP67 PCI Express Bridge	
0563	MCP67 PCI Express Bridge	
0568	MCP78S [GeForce 8200] Memory Controller	
0569	MCP78S [GeForce 8200] PCI Express Bridge	
056a	MCP73 [nForce 630i] USB 2.0 Controller (EHCI)	
056c	MCP73 IDE Controller	
056d	MCP73 PCI Express bridge	
056e	MCP73 PCI Express bridge	
056f	MCP73 PCI Express bridge	
05b1	NF200 PCIe 2.0 switch	
05b8	NF200 PCIe 2.0 switch for GTX 295	
05be	NF200 PCIe 2.0 switch for Quadro Plex S4 / Tesla S870 / Tesla S1070 / Tesla S2050	
05e0	GT200b [GeForce GTX 295]	
05e1	GT200 [GeForce GTX 280]	
05e2	GT200 [GeForce GTX 260]	
05e3	GT200b [GeForce GTX 285]	
05e6	GT200b [GeForce GTX 275]	
05e7	GT200GL [Tesla C1060 / M1060]	
05ea	GT200 [GeForce GTX 260]	
05eb	GT200 [GeForce GTX 295]	
05ed	GT200GL [Quadro Plex 2200 D2]	
05f1	GT200 [GeForce GTX 280]	
05f2	GT200 [GeForce GTX 260]	
05f8	GT200GL [Quadro Plex 2200 S4]	
05f9	GT200GL [Quadro CX]	
05fd	GT200GL [Quadro FX 5800]	
05fe	GT200GL [Quadro FX 4800]	
05ff	GT200GL [Quadro FX 3800]	
0600	G92 [GeForce 8800 GTS 512]	
0601	G92 [GeForce 9800 GT]	
0602	G92 [GeForce 8800 GT]	
0603	G92 [GeForce GT 230 OEM]	
0604	G92 [GeForce 9800 GX2]	
0605	G92 [GeForce 9800 GT]	
0606	G92 [GeForce 8800 GS]	
0607	G92 [GeForce GTS 240]	
0608	G92M [GeForce 9800M GTX]	
0609	G92M [GeForce 8800M GTS]	
060a	G92M [GeForce GTX 280M]	
060b	G92M [GeForce 9800M GT]	
060c	G92M [GeForce 8800M GTX]	
060d	G92 [GeForce 8800 GS]	
060f	G92M [GeForce GTX 285M]	
0610	G92 [GeForce 9600 GSO]	
0611	G92 [GeForce 8800 GT]	
0612	G92 [GeForce 9800 GTX / 9800 GTX+]	
0613	G92 [GeForce 9800 GTX+]	
0614	G92 [GeForce 9800 GT]	
0615	G92 [GeForce GTS 250]	
0617	G92M [GeForce 9800M GTX]	
0618	G92M [GeForce GTX 260M]	
0619	G92GL [Quadro FX 4700 X2]	
061a	G92GL [Quadro FX 3700]	
061b	G92GL [Quadro VX 200]	
061c	G92GLM [Quadro FX 3600M]	
061d	G92GLM [Quadro FX 2800M]	
061e	G92GLM [Quadro FX 3700M]	
061f	G92GLM [Quadro FX 3800M]	
0620	G94 [GeForce 9800 GT]	
0621	G94 [GeForce GT 230]	
0622	G94 [GeForce 9600 GT]	
0623	G94 [GeForce 9600 GS]	
0624	G94 [GeForce 9600 GT Green Edition]	
0625	G94 [GeForce 9600 GSO 512]	
0626	G94 [GeForce GT 130]	
0627	G94 [GeForce GT 140]	
0628	G94M [GeForce 9800M GTS]	
062a	G94M [GeForce 9700M GTS]	
062b	G94M [GeForce 9800M GS]	
062c	G94M [GeForce 9800M GTS]	
062d	G94 [GeForce 9600 GT]	
062e	G94 [GeForce 9600 GT]	
062f	G94 [GeForce 9800 S]	
0630	G94 [GeForce 9600 GT]	
0631	G94M [GeForce GTS 160M]	
0632	G94M [GeForce GTS 150M]	
0633	G94 [GeForce GT 220]	
0635	G94 [GeForce 9600 GSO]	
0637	G94 [GeForce 9600 GT]	
0638	G94GL [Quadro FX 1800]	
063a	G94GLM [Quadro FX 2700M]	
063f	G94 [GeForce 9600 GE]	
0640	G96C [GeForce 9500 GT]	
0641	G96C [GeForce 9400 GT]	
0642	G96 [D9M-10]	
0643	G96 [GeForce 9500 GT]	
0644	G96 [GeForce 9500 GS]	
0645	G96C [GeForce 9500 GS]	
0646	G96C [GeForce GT 120]	
0647	G96CM [GeForce 9600M GT]	
0648	G96CM [GeForce 9600M GS]	
0649	G96CM [GeForce 9600M GT]	
064a	G96M [GeForce 9700M GT]	
064b	G96M [GeForce 9500M G]	
064c	G96CM [GeForce 9650M GT]	
064d		
064e	G96C [GeForce 9600 GSO / 9800 GT]	
0651	G96CM [GeForce G 110M]	
0652	G96CM [GeForce GT 130M]	
0653	G96CM [GeForce GT 120M]	
0654	G96CM [GeForce GT 220M]	
0655	G96 [GeForce GT 120 Mac Edition]	
0656	G96 [GeForce GT 120 Mac Edition]	
0658	G96GL [Quadro FX 380]	
0659	G96CGL [Quadro FX 580]	
065a	G96GLM [Quadro FX 1700M]	
065b	G96C [GeForce 9400 GT]	
065c	G96GLM [Quadro FX 770M]	
065d	G96 [GeForce 9500 GA / 9600 GT / GTS 250]	
065f	G96C [GeForce G210]	
06ac		
06c0	GF100 [GeForce GTX 480]	
06c4	GF100 [GeForce GTX 465]	
06ca	GF100M [GeForce GTX 480M]	
06cb	GF100 [GeForce GTX 480]	
06cd	GF100 [GeForce GTX 470]	
06d0	GF100GL	
06d1	GF100GL [Tesla C2050 / C2070]	
06d2	GF100GL [Tesla M2070]	
06d8	GF100GL [Quadro 6000]	
06d9	GF100GL [Quadro 5000]	
06da	GF100GLM [Quadro 5000M]	
06dc	GF100GL [Quadro 6000]	
06dd	GF100GL [Quadro 4000]	
06de	GF100GL [Tesla T20 Processor]	
06df	GF100GL [Tesla M2070-Q]	
06e0	G98 [GeForce 9300 GE]	
06e1	G98 [GeForce 9300 GS]	
06e2	G98 [GeForce 8400]	
06e3	G98 [GeForce 8300 GS]	
06e4	G98 [GeForce 8400 GS Rev. 2]	
06e5	G98M [GeForce 9300M GS]	
06e6	G98 [GeForce G 100]	
06e7	G98 [GeForce 9300 SE]	
06e8	G98M [GeForce 9200M GS]	
06e9	G98M [GeForce 9300M GS]	
06ea	G98M [Quadro NVS 150M]	
06eb	G98M [Quadro NVS 160M]	
06ec	G98M [GeForce G 105M]	
06ed	G98 [GeForce 9600 GT / 9800 GT]	
06ee	G98 [GeForce 9600 GT / 9800 GT / GT 240]	
06ef	G98M [GeForce G 103M]	
06f1	G98M [GeForce G 105M]	
06f8	G98 [Quadro NVS 420]	
06f9	G98GL [Quadro FX 370 LP]	
06fa	G98 [Quadro NVS 450]	
06fb	G98GLM [Quadro FX 370M]	
06fd	G98 [Quadro NVS 295]	
06ff	G98 [HICx16 + Graphics]	
0751	MCP78S [GeForce 8200] Memory Controller	
0752	MCP78S [GeForce 8200] SMBus	
0753	MCP78S [GeForce 8200] Co-Processor	
0754	MCP78S [GeForce 8200] Memory Controller	
0759	MCP78S [GeForce 8200] IDE	
075a	MCP78S [GeForce 8200] PCI Bridge	
075b	MCP78S [GeForce 8200] PCI Express Bridge	
075c	MCP78S [GeForce 8200] LPC Bridge	
075d	MCP78S [GeForce 8200] LPC Bridge	
0760	MCP77 Ethernet	
0761	MCP77 Ethernet	
0762	MCP77 Ethernet	
0763	MCP77 Ethernet	
0774	MCP72XE/MCP72P/MCP78U/MCP78S High Definition Audio	
0778	MCP78S [GeForce 8200] PCI Express Bridge	
077a	MCP78S [GeForce 8200] PCI Bridge	
077b	MCP78S [GeForce 8200] OHCI USB 1.1 Controller	
077c	MCP78S [GeForce 8200] EHCI USB 2.0 Controller	
077d	MCP78S [GeForce 8200] OHCI USB 1.1 Controller	
077e	MCP78S [GeForce 8200] EHCI USB 2.0 Controller	
07c0	MCP73 Host Bridge	
07c1	MCP73 Host Bridge	
07c2	MCP73 Host Bridge	
07c3	MCP73 Host Bridge	
07c5	MCP73 Host Bridge	
07c8	MCP73 Memory Controller	
07cb	nForce 610i/630i memory controller	
07cd	nForce 610i/630i memory controller	
07ce	nForce 610i/630i memory controller	
07cf	nForce 610i/630i memory controller	
07d0	nForce 610i/630i memory controller	
07d1	nForce 610i/630i memory controller	
07d2	nForce 610i/630i memory controller	
07d3	nForce 610i/630i memory controller	
07d6	nForce 610i/630i memory controller	
07d7	MCP73 LPC Bridge	
07d8	MCP73 SMBus	
07d9	MCP73 Memory Controller	
07da	MCP73 Co-processor	
07dc	MCP73 Ethernet	
07dd	MCP73 Ethernet	
07de	MCP73 Ethernet	
07df	MCP73 Ethernet	
07e0	C73 [GeForce 7150 / nForce 630i]	
07e1	C73 [GeForce 7100 / nForce 630i]	
07e2	C73 [GeForce 7050 / nForce 630i]	
07e3	C73 [GeForce 7050 / nForce 610i]	
07e5	C73 [GeForce 7100 / nForce 620i]	
07f0	MCP73 SATA Controller (IDE mode)	
07f4	GeForce 7100/nForce 630i SATA	
07f8	MCP73 SATA RAID Controller	
07fc	MCP73 High Definition Audio	
07fe	MCP73 OHCI USB 1.1 Controller	
0840	C77 [GeForce 8200M]	
0844	C77 [GeForce 9100M G]	
0845	C77 [GeForce 8200M G]	
0846	C77 [GeForce 9200]	
0847	C78 [GeForce 9100]	
0848	C77 [GeForce 8300]	
0849	C77 [GeForce 8200]	
084a	C77 [nForce 730a]	
084b	C77 [GeForce 8200]	
084c	C77 [nForce 780a/980a SLI]	
084d	C77 [nForce 750a SLI]	
084f	C77 [GeForce 8100 / nForce 720a]	
0860	C79 [GeForce 9300]	
0861	C79 [GeForce 9400]	
0862	C79 [GeForce 9400M G]	
0863	C79 [GeForce 9400M]	
0864	C79 [GeForce 9300]	
0865	C79 [GeForce 9300 / ION]	
0866	C79 [GeForce 9400M G]	
0867	C79 [GeForce 9400]	
0868	C79 [nForce 760i SLI]	
0869	MCP7A [GeForce 9400]	
086a	C79 [GeForce 9400]	
086c	C79 [GeForce 9300 / nForce 730i]	
086d	C79 [GeForce 9200]	
086e	C79 [GeForce 9100M G]	
086f	MCP79 [GeForce 8200M G]	
0870	C79 [GeForce 9400M]	
0871	C79 [GeForce 9200]	
0872	C79 [GeForce G102M]	
0873	C79 [GeForce G102M]	
0874	C79 [ION]	
0876	C79 [GeForce 9400M / ION]	
087a	C79 [GeForce 9400]	
087d	C79 [ION]	
087e	C79 [ION LE]	
087f	C79 [ION LE]	
08a0	MCP89 [GeForce 320M]	
08a2	MCP89 [GeForce 320M]	
08a3	MCP89 [GeForce 320M]	
08a4	MCP89 [GeForce 320M]	
08a5	MCP89 [GeForce 320M]	
0a20	GT216 [GeForce GT 220]	
0a21	GT216M [GeForce GT 330M]	
0a22	GT216 [GeForce 315]	
0a23	GT216 [GeForce 210]	
0a24	GT216 [GeForce 405]	
0a26	GT216 [GeForce 405]	
0a27	GT216 [GeForce 405]	
0a28	GT216M [GeForce GT 230M]	
0a29	GT216M [GeForce GT 330M]	
0a2a	GT216M [GeForce GT 230M]	
0a2b	GT216M [GeForce GT 330M]	
0a2c	GT216M [NVS 5100M]	
0a2d	GT216M [GeForce GT 320M]	
0a30	GT216 [GeForce 505]	
0a32	GT216 [GeForce GT 415]	
0a34	GT216M [GeForce GT 240M]	
0a35	GT216M [GeForce GT 325M]	
0a38	GT216GL [Quadro 400]	
0a3c	GT216GLM [Quadro FX 880M]	
0a60	GT218 [GeForce G210]	
0a62	GT218 [GeForce 205]	
0a63	GT218 [GeForce 310]	
0a64	GT218 [ION]	
0a65	GT218 [GeForce 210]	
0a66	GT218 [GeForce 310]	
0a67	GT218 [GeForce 315]	
0a68	GT218M [GeForce G 105M]	
0a69	GT218M [GeForce G 105M]	
0a6a	GT218M [NVS 2100M]	
0a6c	GT218M [NVS 3100M]	
0a6e	GT218M [GeForce 305M]	
0a6f	GT218M [ION]	
0a70	GT218M [GeForce 310M]	
0a71	GT218M [GeForce 305M]	
0a72	GT218M [GeForce 310M]	
0a73	GT218M [GeForce 305M]	
0a74	GT218M [GeForce G210M]	
0a75	GT218M [GeForce 310M]	
0a76	GT218M [ION 2]	
0a78	GT218GL [Quadro FX 380 LP]	
0a7a	GT218M [GeForce 315M]	
0a7b	GT218 [GeForce 505]	
0a7c	GT218GLM [Quadro FX 380M]	
0a80	MCP79 Host Bridge	
0a81	MCP79 Host Bridge	
0a82	MCP79 Host Bridge	
0a83	MCP79 Host Bridge	
0a84	MCP79 Host Bridge	
0a85	MCP79 Host Bridge	
0a86	MCP79 Host Bridge	
0a87	MCP79 Host Bridge	
0a88	MCP79 Memory Controller	
0a89	MCP79 Memory Controller	
0a98	MCP79 Memory Controller	
0aa0	MCP79 PCI Express Bridge	
0aa2	MCP79 SMBus	
0aa3	MCP79 Co-processor	
0aa4	MCP79 Memory Controller	
0aa5	MCP79 OHCI USB 1.1 Controller	
0aa6	MCP79 EHCI USB 2.0 Controller	
0aa7	MCP79 OHCI USB 1.1 Controller	
0aa8	MCP79 OHCI USB 1.1 Controller	
0aa9	MCP79 EHCI USB 2.0 Controller	
0aaa	MCP79 EHCI USB 2.0 Controller	
0aab	MCP79 PCI Bridge	
0aac	MCP79 LPC Bridge	
0aad	MCP79 LPC Bridge	
0aae	MCP79 LPC Bridge	
0aaf	MCP79 LPC Bridge	
0ab0	MCP79 Ethernet	
0ab1	MCP79 Ethernet	
0ab2	MCP79 Ethernet	
0ab3	MCP79 Ethernet	
0ab4	MCP79 SATA Controller	
0ab5	MCP79 SATA Controller	
0ab6	MCP79 SATA Controller	
0ab7	MCP79 SATA Controller	
0ab8	MCP79 AHCI Controller	
0ab9	MCP79 AHCI Controller	
0aba	MCP79 AHCI Controller	
0abb	MCP79 AHCI Controller	
0abc	MCP79 RAID Controller	
0abd	MCP79 RAID Controller	
0abe	MCP79 RAID Controller	
0abf	MCP79 RAID Controller	
0ac0	MCP79 High Definition Audio	
0ac1	MCP79 High Definition Audio	
0ac2	MCP79 High Definition Audio	
0ac3	MCP79 High Definition Audio	
0ac4	MCP79 PCI Express Bridge	
0ac5	MCP79 PCI Express Bridge	
0ac6	MCP79 PCI Express Bridge	
0ac7	MCP79 PCI Express Bridge	
0ac8	MCP79 PCI Express Bridge	
0ad0	MCP78S [GeForce 8200] SATA Controller (non-AHCI mode)	
0ad4	MCP78S [GeForce 8200] AHCI Controller	
0ad8	MCP78S [GeForce 8200] SATA Controller (RAID mode)	
0be2	GT216 HDMI Audio Controller	
0be3	High Definition Audio Controller	
0be4	High Definition Audio Controller	
0be5	GF100 High Definition Audio Controller	
0be9	GF106 High Definition Audio Controller	
0bea	GF108 High Definition Audio Controller	
0beb	GF104 High Definition Audio Controller	
0bee	GF116 High Definition Audio Controller	
0bf0	Tegra2 PCIe x4 Bridge	
0bf1	Tegra2 PCIe x2 Bridge	
0ca0	GT215 [GeForce GT 330]	
0ca2	GT215 [GeForce GT 320]	
0ca3	GT215 [GeForce GT 240]	
0ca4	GT215 [GeForce GT 340]	
0ca5	GT215 [GeForce GT 220]	
0ca7	GT215 [GeForce GT 330]	
0ca8	GT215M [GeForce GTS 260M]	
0ca9	GT215M [GeForce GTS 250M]	
0cac	GT215 [GeForce GT 220/315]	
0caf	GT215M [GeForce GT 335M]	
0cb0	GT215M [GeForce GTS 350M]	
0cb1	GT215M [GeForce GTS 360M]	
0cbc	GT215GLM [Quadro FX 1800M]	
0d60	MCP89 HOST Bridge	
0d68	MCP89 Memory Controller	
0d69	MCP89 Memory Controller	
0d76	MCP89 PCI Express Bridge	
0d79	MCP89 SMBus	
0d7a	MCP89 Co-Processor	
0d7b	MCP89 Memory Controller	
0d7d	MCP89 Ethernet	
0d80	MCP89 LPC Bridge	
0d85	MCP89 SATA Controller	
0d88	MCP89 SATA Controller (AHCI mode)	
0d89	MCP89 SATA Controller (AHCI mode)	
0d8d	MCP89 SATA Controller (RAID mode)	
0d94	MCP89 High Definition Audio	
0d9c	MCP89 OHCI USB 1.1 Controller	
0d9d	MCP89 EHCI USB 2.0 Controller	
0dc0	GF106 [GeForce GT 440]	
0dc4	GF106 [GeForce GTS 450]	
0dc5	GF106 [GeForce GTS 450 OEM]	
0dc6	GF106 [GeForce GTS 450 OEM]	
0dcd	GF106M [GeForce GT 555M]	
0dce	GF106M [GeForce GT 555M]	
0dd1	GF106M [GeForce GTX 460M]	
0dd2	GF106M [GeForce GT 445M]	
0dd3	GF106M [GeForce GT 435M]	
0dd6	GF106M [GeForce GT 550M]	
0dd8	GF106GL [Quadro 2000]	
0dda	GF106GLM [Quadro 2000M]	
0de0	GF108 [GeForce GT 440]	
0de1	GF108 [GeForce GT 430]	
0de2	GF108 [GeForce GT 420]	
0de3	GF108M [GeForce GT 635M]	
0de4	GF108 [GeForce GT 520]	
0de5	GF108 [GeForce GT 530]	
0de7	GF108 [GeForce GT 610]	
0de8	GF108M [GeForce GT 620M]	
0de9	GF108M [GeForce GT 620M/630M/635M/640M LE]	
0dea	GF108M [GeForce 610M]	
0deb	GF108M [GeForce GT 555M]	
0dec	GF108M [GeForce GT 525M]	
0ded	GF108M [GeForce GT 520M]	
0dee	GF108M [GeForce GT 415M]	
0def	GF108M [NVS 5400M]	
0df0	GF108M [GeForce GT 425M]	
0df1	GF108M [GeForce GT 420M]	
0df2	GF108M [GeForce GT 435M]	
0df3	GF108M [GeForce GT 420M]	
0df4	GF108M [GeForce GT 540M]	
0df5	GF108M [GeForce GT 525M]	
0df6	GF108M [GeForce GT 550M]	
0df7	GF108M [GeForce GT 520M]	
0df8	GF108GL [Quadro 600]	
0df9	GF108GLM [Quadro 500M]	
0dfa	GF108GLM [Quadro 1000M]	
0dfc	GF108GLM [NVS 5200M]	
0e08	GF119 HDMI Audio Controller	
0e09	GF110 High Definition Audio Controller	
0e0a	GK104 HDMI Audio Controller	
0e0b	GK106 HDMI Audio Controller	
0e0c	GF114 HDMI Audio Controller	
0e0f	GK208 HDMI/DP Audio Controller	
0e12	TegraK1 PCIe x4 Bridge	
0e13	TegraK1 PCIe x1 Bridge	
0e1a	GK110 High Definition Audio Controller	
0e1b	GK107 HDMI Audio Controller	
0e1c	Tegra3+ PCIe x4 Bridge	
0e1d	Tegra3+ PCIe x2 Bridge	
0e22	GF104 [GeForce GTX 460]	
0e23	GF104 [GeForce GTX 460 SE]	
0e24	GF104 [GeForce GTX 460 OEM]	
0e30	GF104M [GeForce GTX 470M]	
0e31	GF104M [GeForce GTX 485M]	
0e3a	GF104GLM [Quadro 3000M]	
0e3b	GF104GLM [Quadro 4000M]	
0f00	GF108 [GeForce GT 630]	
0f01	GF108 [GeForce GT 620]	
0f02	GF108 [GeForce GT 730]	
0f03	GF108 [GeForce GT 610]	
0f06	GF108 [GeForce GT 730]	
0f3d		
0fb0	GM200 High Definition Audio	
0fb8	GP108 High Definition Audio Controller	
0fb9	GP107GL High Definition Audio Controller	
0fba	GM206 High Definition Audio Controller	
0fbb	GM204 High Definition Audio Controller	
0fbc	GM107 High Definition Audio Controller [GeForce 940MX]	
0fc0	GK107 [GeForce GT 640 OEM]	
0fc1	GK107 [GeForce GT 640]	
0fc2	GK107 [GeForce GT 630 OEM]	
0fc5	GK107 [GeForce GT 1030]	
0fc6	GK107 [GeForce GTX 650]	
0fc8	GK107 [GeForce GT 740]	
0fc9	GK107 [GeForce GT 730]	
0fcd	GK107M [GeForce GT 755M]	
0fce	GK107M [GeForce GT 640M LE]	
0fd1	GK107M [GeForce GT 650M]	
0fd2	GK107M [GeForce GT 640M]	
0fd3	GK107M [GeForce GT 640M LE]	
0fd4	GK107M [GeForce GTX 660M]	
0fd5	GK107M [GeForce GT 650M Mac Edition]	
0fd6	GK107M	
0fd8	GK107M [GeForce GT 640M Mac Edition]	
0fd9	GK107M [GeForce GT 645M]	
0fdb	GK107M	
0fdf	GK107M [GeForce GT 740M]	
0fe0	GK107M [GeForce GTX 660M Mac Edition]	
0fe1	GK107M [GeForce GT 730M]	
0fe2	GK107M [GeForce GT 745M]	
0fe3	GK107M [GeForce GT 745M]	
0fe4	GK107M [GeForce GT 750M]	
0fe5	GK107 [GeForce K340 USM]	
0fe6	GK107 [GRID K1 NVS USM]	
0fe7	GK107GL [GRID K100 vGPU]	GRID K1 USM
0fe8	GK107M	
0fe9	GK107M [GeForce GT 750M Mac Edition]	
0fea	GK107M [GeForce GT 755M Mac Edition]	
0fec	GK107M [GeForce 710A]	
0fed	GK107M [GeForce 820M]	
0fee	GK107M [GeForce 810M]	
0fef	GK107GL [GRID K340]	
0ff1	GK107 [NVS 1000]	
0ff2	GK107GL [GRID K1]	
0ff3	GK107GL [Quadro K420]	
0ff5	GK107GL [GRID K1 Tesla USM]	
0ff6	GK107GLM [Quadro K1100M]	
0ff7	GK107GL [GRID K140Q vGPU]	GRID K1 Quadro USM
0ff8	GK107GLM [Quadro K500M]	
0ff9	GK107GL [Quadro K2000D]	
0ffa	GK107GL [Quadro K600]	
0ffb	GK107GLM [Quadro K2000M]	
0ffc	GK107GLM [Quadro K1000M]	
0ffd	GK107 [NVS 510]	
0ffe	GK107GL [Quadro K2000]	
0fff	GK107GL [Quadro 410]	
1001	GK110B [GeForce GTX TITAN Z]	
1003	GK110 [GeForce GTX Titan LE]	
1004	GK110 [GeForce GTX 780]	
1005	GK110 [GeForce GTX TITAN]	
1007	GK110 [GeForce GTX 780 Rev. 2]	
1008	GK110 [GeForce GTX 780 Ti 6GB]	
100a	GK110B [GeForce GTX 780 Ti]	
100c	GK110B [GeForce GTX TITAN Black]	
101e	GK110GL [Tesla K20X]	
101f	GK110GL [Tesla K20]	
1020	GK110GL [Tesla K20X]	
1021	GK110GL [Tesla K20Xm]	
1022	GK110GL [Tesla K20c]	
1023	GK110BGL [Tesla K40m]	
1024	GK180GL [Tesla K40c]	
1026	GK110GL [Tesla K20s]	
1027	GK110BGL [Tesla K40st]	
1028	GK110GL [Tesla K20m]	
1029	GK110BGL [Tesla K40s]	
102a	GK110BGL [Tesla K40t]	
102d	GK210GL [Tesla K80]	
102e	GK110BGL [Tesla K40d]	
102f	GK110BGL [Tesla Stella Solo]	
103a	GK110GL [Quadro K6000]	
103c	GK110GL [Quadro K5200]	
103f	GK110BGL [Tesla Stella SXM]	
1040	GF119 [GeForce GT 520]	
1042	GF119 [GeForce 510]	
1045	GF119	
1048	GF119 [GeForce 605]	
1049	GF119 [GeForce GT 620 OEM]	
104a	GF119 [GeForce GT 610]	
104b	GF119 [GeForce GT 625 OEM]	
104c	GF119 [GeForce GT 705]	
104d	GF119 [GeForce GT 710]	
1050	GF119M [GeForce GT 520M]	
1051	GF119M [GeForce GT 520MX]	
1052	GF119M [GeForce GT 520M]	
1054	GF119M [GeForce 410M]	
1055	GF119M [GeForce 410M]	
1056	GF119M [NVS 4200M]	
1057	GF119M [Quadro NVS 4200M]	
1058	GF119M [GeForce 610M]	
1059	GF119M [GeForce 610M]	
105a	GF119M [GeForce 610M]	
105b	GF119M [GeForce 705M]	
107c	GF119 [NVS 315]	
107d	GF119 [NVS 310]	
1080	GF110 [GeForce GTX 580]	
1081	GF110 [GeForce GTX 570]	
1082	GF110 [GeForce GTX 560 Ti OEM]	
1084	GF110 [GeForce GTX 560 OEM]	
1086	GF110 [GeForce GTX 570 Rev. 2]	
1087	GF110 [GeForce GTX 560 Ti 448 Cores]	
1088	GF110 [GeForce GTX 590]	
1089	GF110 [GeForce GTX 580 Rev. 2]	
108b	GF110 [GeForce GTX 580]	
108e	GF110GL [Tesla C2090]	
1091	GF110GL [Tesla M2090]	
1094	GF110GL [Tesla M2075]	
1096	GF110GL [Tesla C2050 / C2075]	
109a	GF100GLM [Quadro 5010M]	
109b	GF100GL [Quadro 7000]	
10c0	GT218 [GeForce 9300 GS Rev. 2]	
10c3	GT218 [GeForce 8400 GS Rev. 3]	
10c5	GT218 [GeForce 405]	
10d8	GT218 [NVS 300]	
10de		
10ef	GP102 HDMI Audio Controller	
10f0	GP104 High Definition Audio Controller	
10f1	GP106 High Definition Audio Controller	
10f7	TU102 High Definition Audio Controller	
10f8	TU104 HD Audio Controller	
10f9	TU106 High Definition Audio Controller	
1140	GF117M [GeForce 610M/710M/810M/820M / GT 620M/625M/630M/720M]	
1180	GK104 [GeForce GTX 680]	
1182	GK104 [GeForce GTX 760 Ti]	
1183	GK104 [GeForce GTX 660 Ti]	
1184	GK104 [GeForce GTX 770]	
1185	GK104 [GeForce GTX 660 OEM]	
1186	GK104 [GeForce GTX 660 Ti]	
1187	GK104 [GeForce GTX 760]	
1188	GK104 [GeForce GTX 690]	
1189	GK104 [GeForce GTX 670]	
118a	GK104GL [GRID K520]	
118b	GK104GL [GRID K2 GeForce USM]	
118c	GK104 [GRID K2 NVS USM]	
118d	GK104GL [GRID K200 vGPU]	GRID K2 USM
118e	GK104 [GeForce GTX 760 OEM]	
118f	GK104GL [Tesla K10]	
1191	GK104 [GeForce GTX 760 Rev. 2]	
1193	GK104 [GeForce GTX 760 Ti OEM]	
1194	GK104GL [Tesla K8]	
1195	GK104 [GeForce GTX 660 Rev. 2]	
1198	GK104M [GeForce GTX 880M]	
1199	GK104M [GeForce GTX 870M]	
119a	GK104M [GeForce GTX 860M]	
119d	GK104M [GeForce GTX 775M Mac Edition]	
119e	GK104M [GeForce GTX 780M Mac Edition]	
119f	GK104M [GeForce GTX 780M]	
11a0	GK104M [GeForce GTX 680M]	
11a1	GK104M [GeForce GTX 670MX]	
11a2	GK104M [GeForce GTX 675MX Mac Edition]	
11a3	GK104M [GeForce GTX 680MX]	
11a7	GK104M [GeForce GTX 675MX]	
11a9	GK104M [GeForce GTX 870M]	
11ac		
11af	GK104GLM [GRID IceCube]	
11b0	GK104GL [GRID K240Q / K260Q vGPU]	
11b1	GK104GL [GRID K2 Tesla USM]	
11b4	GK104GL [Quadro K4200]	
11b6	GK104GLM [Quadro K3100M]	
11b7	GK104GLM [Quadro K4100M]	
11b8	GK104GLM [Quadro K5100M]	
11b9	GK104GLM	
11ba	GK104GL [Quadro K5000]	
11bb	GK104GL [Quadro 4100]	
11bc	GK104GLM [Quadro K5000M]	
11bd	GK104GLM [Quadro K4000M]	
11be	GK104GLM [Quadro K3000M]	
11bf	GK104GL [GRID K2]	
11c0	GK106 [GeForce GTX 660]	
11c2	GK106 [GeForce GTX 650 Ti Boost]	
11c3	GK106 [GeForce GTX 650 Ti OEM]	
11c4	GK106 [GeForce GTX 645 OEM]	
11c5	GK106 [GeForce GT 740]	
11c6	GK106 [GeForce GTX 650 Ti]	
11c7	GK106 [GeForce GTX 750 Ti]	
11c8	GK106 [GeForce GTX 650 OEM]	
11cb	GK106 [GeForce GT 740]	
11e0	GK106M [GeForce GTX 770M]	
11e1	GK106M [GeForce GTX 765M]	
11e2	GK106M [GeForce GTX 765M]	
11e3	GK106M [GeForce GTX 760M]	
11e7	GK106M	
11fa	GK106GL [Quadro K4000]	
11fc	GK106GLM [Quadro K2100M]	
1200	GF114 [GeForce GTX 560 Ti]	
1201	GF114 [GeForce GTX 560]	
1202	GF114 [GeForce GTX 560 Ti OEM]	
1203	GF114 [GeForce GTX 460 SE v2]	
1205	GF114 [GeForce GTX 460 v2]	
1206	GF114 [GeForce GTX 555]	
1207	GF114 [GeForce GT 645 OEM]	
1208	GF114 [GeForce GTX 560 SE]	
1210	GF114M [GeForce GTX 570M]	
1211	GF114M [GeForce GTX 580M]	
1212	GF114M [GeForce GTX 675M]	
1213	GF114M [GeForce GTX 670M]	
1241	GF116 [GeForce GT 545 OEM]	
1243	GF116 [GeForce GT 545]	
1244	GF116 [GeForce GTX 550 Ti]	
1245	GF116 [GeForce GTS 450 Rev. 2]	
1246	GF116M [GeForce GT 550M]	
1247	GF116M [GeForce GT 555M/635M]	
1248	GF116M [GeForce GT 555M/635M]	
1249	GF116 [GeForce GTS 450 Rev. 3]	
124b	GF116 [GeForce GT 640 OEM]	
124d	GF116M [GeForce GT 555M/635M]	
1251	GF116M [GeForce GT 560M]	
1280	GK208 [GeForce GT 635]	
1281	GK208 [GeForce GT 710]	
1282	GK208 [GeForce GT 640 Rev. 2]	
1284	GK208 [GeForce GT 630 Rev. 2]	
1286	GK208 [GeForce GT 720]	
1287	GK208B [GeForce GT 730]	
1288	GK208B [GeForce GT 720]	
1289	GK208 [GeForce GT 710]	
128a	GK208B	
128b	GK208B [GeForce GT 710]	
128c	GK208B	
1290	GK208M [GeForce GT 730M]	
1291	GK208M [GeForce GT 735M]	
1292	GK208M [GeForce GT 740M]	
1293	GK208M [GeForce GT 730M]	
1294	GK208M [GeForce GT 740M]	
1295	GK208M [GeForce 710M]	
1296	GK208M [GeForce 825M]	
1298	GK208M [GeForce GT 720M]	
1299	GK208BM [GeForce 920M]	
129a	GK208BM [GeForce 910M]	
12a0	GK208	
12b9	GK208GLM [Quadro K610M]	
12ba	GK208GLM [Quadro K510M]	
1340	GM108M [GeForce 830M]	
1341	GM108M [GeForce 840M]	
1344	GM108M [GeForce 845M]	
1346	GM108M [GeForce 930M]	
1347	GM108M [GeForce 940M]	
1348	GM108M [GeForce 945M / 945A]	
1349	GM108M [GeForce 930M]	
134b	GM108M [GeForce 940MX]	
134d	GM108M [GeForce 940MX]	
134e	GM108M [GeForce 930MX]	
134f	GM108M [GeForce 920MX]	
137a	GM108GLM [Quadro K620M / Quadro M500M]	
137b	GM108GLM [Quadro M520 Mobile]	
137d	GM108M [GeForce 940A]	
1380	GM107 [GeForce GTX 750 Ti]	
1381	GM107 [GeForce GTX 750]	
1382	GM107 [GeForce GTX 745]	
1389	GM107GL [GRID M30]	
1390	GM107M [GeForce 845M]	
1391	GM107M [GeForce GTX 850M]	
1392	GM107M [GeForce GTX 860M]	
1393	GM107M [GeForce 840M]	
1398	GM107M [GeForce 845M]	
1399	GM107M [GeForce 945M]	
139a	GM107M [GeForce GTX 950M]	
139b	GM107M [GeForce GTX 960M]	
139c	GM107M [GeForce 940M]	
139d	GM107M [GeForce GTX 750 Ti]	
13b0	GM107GLM [Quadro M2000M]	
13b1	GM107GLM [Quadro M1000M]	
13b2	GM107GLM [Quadro M600M]	
13b3	GM107GLM [Quadro K2200M]	
13b4	GM107GLM [Quadro M620 Mobile]	
13b6	GM107GLM [Quadro M1200 Mobile]	
13b9	GM107GL [NVS 810]	
13ba	GM107GL [Quadro K2200]	
13bb	GM107GL [Quadro K620]	
13bc	GM107GL [Quadro K1200]	
13bd	GM107GL [Tesla M10]	
13c0	GM204 [GeForce GTX 980]	
13c1	GM204	
13c2	GM204 [GeForce GTX 970]	
13c3	GM204	
13d7	GM204M [GeForce GTX 980M]	
13d8	GM204M [GeForce GTX 970M]	
13d9	GM204M [GeForce GTX 965M]	
13da	GM204M [GeForce GTX 980 Mobile]	
13e7	GM204GL [GeForce GTX 980 Engineering Sample]	
13f0	GM204GL [Quadro M5000]	
13f1	GM204GL [Quadro M4000]	
13f2	GM204GL [Tesla M60]	
13f3	GM204GL [Tesla M6]	
13f8	GM204GLM [Quadro M5000M / M5000 SE]	
13f9	GM204GLM [Quadro M4000M]	
13fa	GM204GLM [Quadro M3000M]	
13fb	GM204GLM [Quadro M5500]	
1401	GM206 [GeForce GTX 960]	
1402	GM206 [GeForce GTX 950]	
1404	GM206 [GeForce GTX 960 FAKE]	
1406	GM206 [GeForce GTX 960 OEM]	
1407	GM206 [GeForce GTX 750 v2]	
1427	GM206M [GeForce GTX 965M]	
1430	GM206GL [Quadro M2000]	
1431	GM206GL [Tesla M4]	
1436	GM206GLM [Quadro M2200 Mobile]	
15f0	GP100GL [Quadro GP100]	
15f1	GP100GL	
15f7	GP100GL [Tesla P100 PCIe 12GB]	
15f8	GP100GL [Tesla P100 PCIe 16GB]	
15f9	GP100GL [Tesla P100 SXM2 16GB]	
1617	GM204M [GeForce GTX 980M]	
1618	GM204M [GeForce GTX 970M]	
1619	GM204M [GeForce GTX 965M]	
161a	GM204M [GeForce GTX 980 Mobile]	
1667	GM204M [GeForce GTX 965M]	
1725	GP100	
172e	GP100	
172f	GP100	
174d	GM108M [GeForce MX130]	
174e	GM108M [GeForce MX110]	
1789	GM107GL [GRID M3-3020]	
179c	GM107 [GeForce 940MX]	
17c2	GM200 [GeForce GTX TITAN X]	
17c8	GM200 [GeForce GTX 980 Ti]	
17f0	GM200GL [Quadro M6000]	
17f1	GM200GL [Quadro M6000 24GB]	
17fd	GM200GL [Tesla M40]	
1ad0	Tegra PCIe x8 Endpoint	
1ad1	Tegra PCIe x4/x8 Endpoint/Root Complex	
1ad2	Tegra PCIe x1 Root Complex	
1ad3	Xavier SATA Controller	
1ad6	TU102 USB 3.1 Host Controller	
1ad7	TU102 USB Type-C UCSI Controller	
1ad8	TU104 USB 3.1 Host Controller	
1ad9	TU104 USB Type-C UCSI Controller	
1ada	TU106 USB 3.1 Host Controller	
1adb	TU106 USB Type-C UCSI Controller	
1aeb	TU116 High Definition Audio Controller	
1aec	TU116 USB 3.1 Host Controller	
1aed	TU116 USB Type-C UCSI Controller	
1aef	GA102 High Definition Audio Controller	
1b00	GP102 [TITAN X]	
1b01	GP102 [GeForce GTX 1080 Ti 10GB]	
1b02	GP102 [TITAN Xp]	
1b04	GP102	
1b06	GP102 [GeForce GTX 1080 Ti]	
1b07	GP102 [P102-100]	
1b30	GP102GL [Quadro P6000]	
1b38	GP102GL [Tesla P40]	
1b39	GP102GL [Tesla P10]	
1b70	GP102GL	
1b78	GP102GL	
1b80	GP104 [GeForce GTX 1080]	
1b81	GP104 [GeForce GTX 1070]	
1b82	GP104 [GeForce GTX 1070 Ti]	
1b83	GP104 [GeForce GTX 1060 6GB]	
1b84	GP104 [GeForce GTX 1060 3GB]	
1b87	GP104 [P104-100]	
1ba0	GP104M [GeForce GTX 1080 Mobile]	
1ba1	GP104M [GeForce GTX 1070 Mobile]	
1ba2	GP104M [GeForce GTX 1070 Mobile]	
1ba9	GP104M	
1baa	GP104M	
1bad	GP104 [GeForce GTX 1070 Engineering Sample]	
1bb0	GP104GL [Quadro P5000]	
1bb1	GP104GL [Quadro P4000]	
1bb3	GP104GL [Tesla P4]	
1bb4	GP104GL [Tesla P6]	
1bb5	GP104GLM [Quadro P5200 Mobile]	
1bb6	GP104GLM [Quadro P5000 Mobile]	
1bb7	GP104GLM [Quadro P4000 Mobile]	
1bb8	GP104GLM [Quadro P3000 Mobile]	
1bb9	GP104GLM [Quadro P4200 Mobile]	
1bbb	GP104GLM [Quadro P3200 Mobile]	
1bc7	GP104 [P104-101]	
1be0	GP104BM [GeForce GTX 1080 Mobile]	
1be1	GP104BM [GeForce GTX 1070 Mobile]	
1c00	GP106	
1c01	GP106	
1c02	GP106 [GeForce GTX 1060 3GB]	
1c03	GP106 [GeForce GTX 1060 6GB]	
1c04	GP106 [GeForce GTX 1060 5GB]	
1c06	GP106 [GeForce GTX 1060 6GB Rev. 2]	
1c07	GP106 [P106-100]	
1c09	GP106 [P106-090]	
1c20	GP106M [GeForce GTX 1060 Mobile]	
1c21	GP106M [GeForce GTX 1050 Ti Mobile]	
1c22	GP106M [GeForce GTX 1050 Mobile]	
1c23	GP106M [GeForce GTX 1060 Mobile Rev. 2]	
1c2d	GP106M	
1c30	GP106GL [Quadro P2000]	
1c31	GP106GL [Quadro P2200]	
1c35	GP106M [Quadro P2000 Mobile]	
1c36	GP106 [P106M]	
1c60	GP106BM [GeForce GTX 1060 Mobile 6GB]	
1c61	GP106BM [GeForce GTX 1050 Ti Mobile]	
1c62	GP106BM [GeForce GTX 1050 Mobile]	
1c70	GP106GL	
1c80		
1c81	GP107 [GeForce GTX 1050]	
1c82	GP107 [GeForce GTX 1050 Ti]	
1c83	GP107 [GeForce GTX 1050 3GB]	
1c8c	GP107M [GeForce GTX 1050 Ti Mobile]	
1c8d	GP107M [GeForce GTX 1050 Mobile]	
1c8e	GP107M	
1c8f	GP107M [GeForce GTX 1050 Ti Max-Q]	
1c90	GP107M [GeForce MX150]	
1c91	GP107M [GeForce GTX 1050 3 GB Max-Q]	
1c92	GP107M [GeForce GTX 1050 Mobile]	
1c94	GP107M [GeForce MX350]	
1c96	GP107M [GeForce MX350]	
1ca7	GP107GL	
1ca8	GP107GL	
1caa	GP107GL	
1cb1	GP107GL [Quadro P1000]	
1cb2	GP107GL [Quadro P600]	
1cb3	GP107GL [Quadro P400]	
1cb6	GP107GL [Quadro P620]	
1cba	GP107GLM [Quadro P2000 Mobile]	
1cbb	GP107GLM [Quadro P1000 Mobile]	
1cbc	GP107GLM [Quadro P600 Mobile]	
1cbd	GP107GLM [Quadro P620]	
1ccc	GP107BM [GeForce GTX 1050 Ti Mobile]	
1ccd	GP107BM [GeForce GTX 1050 Mobile]	
1cfa	GP107GL [Quadro P2000]	
1cfb	GP107GL [Quadro P1000]	
1d01	GP108 [GeForce GT 1030]	
1d02	GP108 [GeForce GT 1010]	
1d10	GP108M [GeForce MX150]	
1d11	GP108M [GeForce MX230]	
1d12	GP108M [GeForce MX150]	
1d13	GP108M [GeForce MX250]	
1d16	GP108M [GeForce MX330]	
1d33	GP108GLM [Quadro P500 Mobile]	
1d34	GP108GLM [Quadro P520]	
1d52	GP108BM [GeForce MX250]	
1d56	GP108BM [GeForce MX330]	
1d81	GV100 [TITAN V]	
1db1	GV100GL [Tesla V100 SXM2 16GB]	
1db2	GV100GL [Tesla V100 DGXS 16GB]	
1db3	GV100GL [Tesla V100 FHHL 16GB]	
1db4	GV100GL [Tesla V100 PCIe 16GB]	
1db5	GV100GL [Tesla V100 SXM2 32GB]	
1db6	GV100GL [Tesla V100 PCIe 32GB]	
1db7	GV100GL [Tesla V100 DGXS 32GB]	
1db8	GV100GL [Tesla V100 SXM3 32GB]	
1dba	GV100GL [Quadro GV100]	
1df0	GV100GL [Tesla PG500-216]	
1df2	GV100GL [Tesla PG503-216]	
1df5	GV100GL [Tesla V100 SXM2 16GB]	
1df6	GV100GL [Tesla V100S PCIe 32GB]	
1e02	TU102 [TITAN RTX]	
1e04	TU102 [GeForce RTX 2080 Ti]	
1e07	TU102 [GeForce RTX 2080 Ti Rev. A]	
1e09	TU102 [CMP 50HX]	
1e2d	TU102 [GeForce RTX 2080 Ti Engineering Sample]	
1e2e	TU102 [GeForce RTX 2080 Ti 12GB Engineering Sample]	
1e30	TU102GL [Quadro RTX 6000/8000]	
1e36	TU102GL [Quadro RTX 6000]	
1e37	TU102GL [GRID RTX T10-4/T10-8/T10-16]	
1e38	TU102GL	
1e3c	TU102GL	
1e3d	TU102GL	
1e3e	TU102GL	
1e78	TU102GL [Quadro RTX 6000/8000]	
1e81	TU104 [GeForce RTX 2080 SUPER]	
1e82	TU104 [GeForce RTX 2080]	
1e84	TU104 [GeForce RTX 2070 SUPER]	
1e87	TU104 [GeForce RTX 2080 Rev. A]	
1e89	TU104 [GeForce RTX 2060]	
1e90	TU104M [GeForce RTX 2080 Mobile]	
1e91	TU104M [GeForce RTX 2070 SUPER Mobile / Max-Q]	
1e93	TU104M [GeForce RTX 2080 SUPER Mobile / Max-Q]	
1eab	TU104M	
1eae	TU104M	
1eb0	TU104GL [Quadro RTX 5000]	
1eb1	TU104GL [Quadro RTX 4000]	
1eb4	TU104GL [T4G]	
1eb5	TU104GLM [Quadro RTX 5000 Mobile / Max-Q]	
1eb6	TU104GLM [Quadro RTX 4000 Mobile / Max-Q]	
1eb8	TU104GL [Tesla T4]	
1eb9	TU104GL	
1ebe	TU104GL	
1ec2	TU104 [GeForce RTX 2070 SUPER]	
1ec7	TU104 [GeForce RTX 2070 SUPER]	
1ed0	TU104BM [GeForce RTX 2080 Mobile]	
1ed1	TU104BM [GeForce RTX 2070 SUPER Mobile / Max-Q]	
1ed3	TU104BM [GeForce RTX 2080 SUPER Mobile / Max-Q]	
1ef5	TU104GLM [Quadro RTX 5000 Mobile Refresh]	
1f02	TU106 [GeForce RTX 2070]	
1f03	TU106 [GeForce RTX 2060 12GB]	
1f04	TU106	
1f06	TU106 [GeForce RTX 2060 SUPER]	
1f07	TU106 [GeForce RTX 2070 Rev. A]	
1f08	TU106 [GeForce RTX 2060 Rev. A]	
1f09	TU106 [GeForce GTX 1660 SUPER]	
1f0a	TU106 [GeForce GTX 1650]	
1f0b	TU106 [CMP 40HX]	
1f10	TU106M [GeForce RTX 2070 Mobile]	
1f11	TU106M [GeForce RTX 2060 Mobile]	
1f12	TU106M [GeForce RTX 2060 Max-Q]	
1f14	TU106M [GeForce RTX 2070 Mobile / Max-Q Refresh]	
1f15	TU106M [GeForce RTX 2060 Mobile]	
1f2e	TU106M	
1f36	TU106GLM [Quadro RTX 3000 Mobile / Max-Q]	
1f42	TU106 [GeForce RTX 2060 SUPER]	
1f47	TU106 [GeForce RTX 2060 SUPER]	
1f50	TU106BM [GeForce RTX 2070 Mobile / Max-Q]	
1f51	TU106BM [GeForce RTX 2060 Mobile]	
1f54	TU106BM [GeForce RTX 2070 Mobile]	
1f55	TU106BM [GeForce RTX 2060 Mobile]	
1f76	TU106GLM [Quadro RTX 3000 Mobile Refresh]	
1f81	TU117	
1f82	TU117 [GeForce GTX 1650]	
1f83	TU117 [GeForce GTX 1630]	
1f91	TU117M [GeForce GTX 1650 Mobile / Max-Q]	
1f92	TU117M [GeForce GTX 1650 Mobile]	
1f94	TU117M [GeForce GTX 1650 Mobile]	
1f95	TU117M [GeForce GTX 1650 Ti Mobile]	
1f96	TU117M [GeForce GTX 1650 Mobile / Max-Q]	
1f97	TU117M [GeForce MX450]	
1f98	TU117M [GeForce MX450]	
1f99	TU117M	
1f9c	TU117M [GeForce MX450]	
1f9d	TU117M [GeForce GTX 1650 Mobile / Max-Q]	
1f9f	TU117M [GeForce MX550]	via Lenovo 496.90
1fa0	TU117M [GeForce MX550]	
1fae	TU117GL	
1fb0	TU117GLM [Quadro T1000 Mobile]	
1fb1	TU117GL [T600]	
1fb2	TU117GLM [Quadro T400 Mobile]	
1fb6	TU117GLM [T600 Laptop GPU]	
1fb7	TU117GLM [T550 Laptop GPU]	
1fb8	TU117GLM [Quadro T2000 Mobile / Max-Q]	
1fb9	TU117GLM [Quadro T1000 Mobile]	
1fba	TU117GLM [T600 Mobile]	
1fbb	TU117GLM [Quadro T500 Mobile]	
1fbc	TU117GLM [T1200 Laptop GPU]	
1fbf	TU117GL	
1fd9	TU117BM [GeForce GTX 1650 Mobile Refresh]	
1fdd	TU117BM [GeForce GTX 1650 Mobile Refresh]	
1ff0	TU117GL [T1000 8GB]	
1ff2	TU117GL [T400 4GB]	
1ff9	TU117GLM [Quadro T1000 Mobile]	
2082	GA100 [CMP 170HX]	
20b0	GA100 [A100 SXM4 40GB]	
20b1	GA100 [A100 PCIe 40GB]	
20b2	GA100 [A100 SXM4 80GB]	
20b3	GA100 [PG506-242/243]	20B3 14A7 10DE PG506-242 / 20B3 14A8 10DE PG506-243
20b5	GA100 [A100 PCIe 80GB]	
20b6	GA100GL [PG506-232]	
20b7	GA100GL [A30 PCIe]	
20b8	GA100 [A100X]	
20b9	GA100 [A30X]	
20bb	GA100 [DRIVE A100 PROD]	
20be	GA100 [GRID A100A]	
20bf	GA100 [GRID A100B]	
20c2	GA100 [CMP 170HX]	
20f0	GA100 [A100-PG506-207]	
20f1	GA100 [A100 PCIe 40GB]	
20f2	GA100 [A100-PG506-217]	
2182	TU116 [GeForce GTX 1660 Ti]	
2183	TU116	
2184	TU116 [GeForce GTX 1660]	
2187	TU116 [GeForce GTX 1650 SUPER]	
2188	TU116 [GeForce GTX 1650]	
2189	TU116 [CMP 30HX]	
2191	TU116M [GeForce GTX 1660 Ti Mobile]	
2192	TU116M [GeForce GTX 1650 Ti Mobile]	
21ae	TU116GL	
21bf	TU116GL	
21c2	TU116	
21c3		
21c4	TU116 [GeForce GTX 1660 SUPER]	
21d1	TU116BM [GeForce GTX 1660 Ti Mobile]	
2200	GA102	
2203	GA102 [GeForce RTX 3090 Ti]	
2204	GA102 [GeForce RTX 3090]	
2205	GA102 [GeForce RTX 3080 Ti 20GB]	
2206	GA102 [GeForce RTX 3080]	
2208	GA102 [GeForce RTX 3080 Ti]	
220a	GA102 [GeForce RTX 3080 12GB]	
220d	GA102 [CMP 90HX]	
2216	GA102 [GeForce RTX 3080 Lite Hash Rate]	
222b	GA102 [GeForce RTX 3090 Engineering Sample]	
222f	GA102 [GeForce RTX 3080 11GB / 12GB Engineering Sample]	
2230	GA102GL [RTX A6000]	
2231	GA102GL [RTX A5000]	
2232	GA102GL [RTX A4500]	
2233	GA102GL [RTX A5500]	
2235	GA102GL [A40]	
2236	GA102GL [A10]	
2237	GA102GL [A10G]	
2238	GA102GL [A10M]	
223f	GA102GL	
228b	GA104 High Definition Audio Controller	
228e	GA106 High Definition Audio Controller	
2296	Tegra PCIe Endpoint Virtual Network	
2302	GH100	
2321	GH100	
2336	GH100 [H100 96GB]	
2414	GA103 [GeForce RTX 3060 Ti]	
2420	GA103M [GeForce RTX 3080 Ti Mobile]	
2438	GA103GLM [RTX A5500 Laptop GPU]	
2460	GA103M [GeForce RTX 3080 Ti Laptop GPU]	
2482	GA104 [GeForce RTX 3070 Ti]	
2483	GA104	
2484	GA104 [GeForce RTX 3070]	
2486	GA104 [GeForce RTX 3060 Ti]	
2487	GA104 [GeForce RTX 3060]	
2488	GA104 [GeForce RTX 3070 Lite Hash Rate]	
2489	GA104 [GeForce RTX 3060 Ti Lite Hash Rate]	
248a	GA104 [CMP 70HX]	
249c	GA104M [GeForce RTX 3080 Mobile / Max-Q 8GB/16GB]	
249d	GA104M [GeForce RTX 3070 Mobile / Max-Q]	
249f	GA104M	
24a0	GA104 [Geforce RTX 3070 Ti Laptop GPU]	
24ac	GA104 [GeForce RTX 30x0 Engineering Sample]	
24ad	GA104 [GeForce RTX 3060 Engineering Sample]	
24af	GA104 [GeForce RTX 3070 Engineering Sample]	
24b0	GA104GL [RTX A4000]	
24b1	GA104GL [RTX A4000H]	
24b6	GA104GLM [RTX A5000 Mobile]	
24b7	GA104GLM [RTX A4000 Mobile]	
24b8	GA104GLM [RTX A3000 Mobile]	
24b9	GA104GLM [RTX A3000 12GB Laptop GPU]	
24ba	GA104GLM [RTX A4500 Laptop GPU]	
24bb	GA104GLM [RTX A3000 Laptop GPU]	
24bf	GA104 [GeForce RTX 3070 Engineering Sample]	
24dc	GA104M [GeForce RTX 3080 Mobile / Max-Q 8GB/16GB]	
24dd	GA104M [GeForce RTX 3070 Mobile / Max-Q]	
24e0	GA104M [Geforce RTX 3070 Ti Laptop GPU]	
24fa	GA104 [RTX A4500 Embedded GPU ]	
2501	GA106 [GeForce RTX 3060]	
2503	GA106 [GeForce RTX 3060]	
2504	GA106 [GeForce RTX 3060 Lite Hash Rate]	
2505	GA106	
2507	GA106 [Geforce RTX 3050]	
2508	GA106 [GeForce RTX 3050 OEM]	
2520	GA106M [GeForce RTX 3060 Mobile / Max-Q]	
2523	GA106M [GeForce RTX 3050 Ti Mobile / Max-Q]	
252f	GA106 [GeForce RTX 3060 Engineering Sample]	
2531	GA106 [RTX A2000]	
2560	GA106M [GeForce RTX 3060 Mobile / Max-Q]	
2563	GA106M [GeForce RTX 3050 Ti Mobile / Max-Q]	
2571	GA106 [RTX A2000 12GB]	
2583	GA107 [GeForce RTX 3050]	
25a0	GA107M [GeForce RTX 3050 Ti Mobile]	
25a2	GA107M [GeForce RTX 3050 Mobile]	
25a3	GA107	
25a4	GA107	
25a5	GA107M [GeForce RTX 3050 Mobile]	
25a6	GA107M [GeForce MX570]	
25a7	GA107M [GeForce MX570]	
25a9	GA107M [GeForce RTX 2050]	
25aa	GA107M [GeForce MX570 A]	
25ac	GN20-P0-R-K2	GN20-P0 Refresh
25af	GA107 [GeForce RTX 3050 Engineering Sample]	
25b5	GA107GLM [RTX A4 Mobile]	
25b6	GA107GL [A2 / A16]	A16 - 25B6 10DE 14A9 / A2 - 25B6 10DE 157E
25b8	GA107GLM [RTX A2000 Mobile]	
25b9	GA107GLM [RTX A1000 Laptop GPU]	
25ba	GA107GLM [RTX A2000 8GB Laptop GPU]	
25bb	GA107GLM [RTX A500 Laptop GPU]	
25e0	GA107BM [GeForce RTX 3050 Ti Mobile]	
25e2	GA107BM [GeForce RTX 3050 Mobile]	
25e5	GA107BM [GeForce RTX 3050 Mobile]	
25ec	GN20-P0-R-K2	GN20-P0 Refresh
25f9	GA107 [RTX A1000 Embedded GPU ]	
25fa	GA107 [RTX A2000 Embedded GPU]	
2684	AD102 [GeForce RTX 4090]	
2717	GN21-X11	
2757	GN21-X11	
2785	AD104	
27a0	GN21-X9	
27e0	GN21-X9	
2820	GN21-X6	
2860	GN21-X6	
28a0	GN21-X4	
28a1	GN21-X2	
28e0	GN21-X4	
28e1	GN21-X2
`
