--Create table
CREATE TABLE IF NOT EXISTS BANKS (
  BankID INTEGER PRIMARY KEY,
  BankName CHAR(256) NOT NULL,
  BankFullname TEXT NOT NULL,
  CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  LastUpdate TIMESTAMP DEFAULT CURRENT_TIMESTAMP  
);

--Clean data
DELETE FROM BANKS;

--Insert data
INSERT INTO BANKS (	BankID	,	BankName	,	BankFullName	) VALUES
(	1	,'BCO DO BRASIL S.A.','Banco do Brasil S.A.'),
(	3	,'BCO DA AMAZONIA S.A.','BANCO DA AMAZONIA S.A.'),
(	10	,'CREDICOAMO','CREDICOAMO CREDITO RURAL COOPERATIVA'),
(	11	,'C.SUISSE HEDGING-GRIFFO CV S/A','CREDIT SUISSE HEDGING-GRIFFO CORRETORA DE VALORES S.A'),
(	12	,'BANCO INBURSA','Banco Inbursa S.A.'),
(	15	,'UBS BRASIL CCTVM S.A.','UBS Brasil Corretora de Câmbio, Títulos e Valores Mobiliários S.A.'),
(	16	,'SICOOB CREDITRAN','COOPERATIVA DE CRÉDITO MÚTUO DOS DESPACHANTES DE TRÂNSITO DE SANTA CATARINA E RI'),
(	17	,'BNY MELLON BCO S.A.','BNY Mellon Banco S.A.'),
(	18	,'BCO TRICURY S.A.','Banco Tricury S.A.'),
(	21	,'BCO BANESTES S.A.','BANESTES S.A. BANCO DO ESTADO DO ESPIRITO SANTO'),
(	24	,'BCO BANDEPE S.A.','Banco Bandepe S.A.'),
(	25	,'BCO ALFA S.A.','Banco Alfa S.A.'),
(	29	,'BANCO ITAÚ CONSIGNADO S.A.','Banco Itaú Consignado S.A.'),
(	33	,'BCO SANTANDER (BRASIL) S.A.','BANCO SANTANDER (BRASIL) S.A.'),
(	36	,'BCO BBI S.A.','Banco Bradesco BBI S.A.'),
(	37	,'BCO DO EST. DO PA S.A.','Banco do Estado do Pará S.A.'),
(	40	,'BCO CARGILL S.A.','Banco Cargill S.A.'),
(	41	,'BCO DO ESTADO DO RS S.A.','Banco do Estado do Rio Grande do Sul S.A.'),
(	47	,'BCO DO EST. DE SE S.A.','Banco do Estado de Sergipe S.A.'),
(	60	,'CONFIDENCE CC S.A.','Confidence Corretora de Câmbio S.A.'),
(	62	,'HIPERCARD BM S.A.','Hipercard Banco Múltiplo S.A.'),
(	63	,'BANCO BRADESCARD','Banco Bradescard S.A.'),
(	64	,'GOLDMAN SACHS DO BRASIL BM S.A','GOLDMAN SACHS DO BRASIL BANCO MULTIPLO S.A.'),
(	65	,'BCO ANDBANK S.A.','Banco AndBank (Brasil) S.A.'),
(	66	,'BCO MORGAN STANLEY S.A.','BANCO MORGAN STANLEY S.A.'),
(	69	,'BCO CREFISA S.A.','Banco Crefisa S.A.'),
(	70	,'BRB - BCO DE BRASILIA S.A.','BRB - BANCO DE BRASILIA S.A.'),
(	75	,'BCO ABN AMRO S.A.','Banco ABN Amro S.A.'),
(	76	,'BCO KDB BRASIL S.A.','Banco KDB do Brasil S.A.'),
(	77	,'BANCO INTER','Banco Inter S.A.'),
(	78	,'HAITONG BI DO BRASIL S.A.','Haitong Banco de Investimento do Brasil S.A.'),
(	79	,'BCO ORIGINAL DO AGRO S/A','Banco Original do Agronegócio S.A.'),
(	80	,'B&T CC LTDA.','B&T CORRETORA DE CAMBIO LTDA.'),
(	81	,'BANCOSEGURO S.A.','BancoSeguro S.A.'),
(	82	,'BANCO TOPÁZIO S.A.','BANCO TOPÁZIO S.A.'),
(	83	,'BCO DA CHINA BRASIL S.A.','Banco da China Brasil S.A.'),
(	84	,'UNIPRIME NORTE DO PARANÁ - CC','UNIPRIME NORTE DO PARANÁ - COOPERATIVA DE CRÉDITO LTDA'),
(	85	,'COOP CENTRAL AILOS','Cooperativa Central de Crédito - Ailos'),
(	89	,'CREDISAN CC','CREDISAN COOPERATIVA DE CRÉDITO'),
(	91	,'CCCM UNICRED CENTRAL RS','CENTRAL DE COOPERATIVAS DE ECONOMIA E CRÉDITO MÚTUO DO ESTADO DO RIO GRANDE DO S'),
(	92	,'BRK S.A. CFI','BRK S.A. Crédito, Financiamento e Investimento'),
(	94	,'BANCO FINAXIS','Banco Finaxis S.A.'),
(	95	,'TRAVELEX BANCO DE CÂMBIO S.A.','Travelex Banco de Câmbio S.A.'),
(	96	,'BCO B3 S.A.','Banco B3 S.A.'),
(	97	,'CREDISIS CENTRAL DE COOPERATIVAS DE CRÉDITO LTDA.','Credisis - Central de Cooperativas de Crédito Ltda.'),
(	98	,'CREDIALIANÇA CCR','Credialiança Cooperativa de Crédito Rural'),
(	99	,'UNIPRIME CENTRAL CCC LTDA.','UNIPRIME CENTRAL - CENTRAL INTERESTADUAL DE COOPERATIVAS DE CREDITO LTDA.'),
(	100	,'PLANNER CV S.A.','Planner Corretora de Valores S.A.'),
(	101	,'RENASCENCA DTVM LTDA','RENASCENCA DISTRIBUIDORA DE TÍTULOS E VALORES MOBILIÁRIOS LTDA'),
(	102	,'XP INVESTIMENTOS CCTVM S/A','XP INVESTIMENTOS CORRETORA DE CÂMBIO,TÍTULOS E VALORES MOBILIÁRIOS S/A'),
(	104	,'CAIXA ECONOMICA FEDERAL','CAIXA ECONOMICA FEDERAL'),
(	105	,'LECCA CFI S.A.','Lecca Crédito, Financiamento e Investimento S/A'),
(	107	,'BCO BOCOM BBM S.A.','Banco Bocom BBM S.A.'),
(	108	,'PORTOCRED S.A. - CFI','PORTOCRED S.A. - CREDITO, FINANCIAMENTO E INVESTIMENTO'),
(	111	,'OLIVEIRA TRUST DTVM S.A.','OLIVEIRA TRUST DISTRIBUIDORA DE TÍTULOS E VALORES MOBILIARIOS S.A.'),
(	113	,'MAGLIANO S.A. CCVM','Magliano S.A. Corretora de Cambio e Valores Mobiliarios'),
(	114	,'CENTRAL COOPERATIVA DE CRÉDITO NO ESTADO DO ESPÍRITO SANTO','Central Cooperativa de Crédito no Estado do Espírito Santo - CECOOP'),
(	117	,'ADVANCED CC LTDA','ADVANCED CORRETORA DE CÂMBIO LTDA'),
(	119	,'BCO WESTERN UNION','Banco Western Union do Brasil S.A.'),
(	120	,'BCO RODOBENS S.A.','BANCO RODOBENS S.A.'),
(	121	,'BCO AGIBANK S.A.','Banco Agibank S.A.'),
(	122	,'BCO BRADESCO BERJ S.A.','Banco Bradesco BERJ S.A.'),
(	124	,'BCO WOORI BANK DO BRASIL S.A.','Banco Woori Bank do Brasil S.A.'),
(	125	,'PLURAL BCO BM','Plural S.A. Banco Múltiplo'),
(	126	,'BR PARTNERS BI','BR Partners Banco de Investimento S.A.'),
(	127	,'CODEPE CVC S.A.','Codepe Corretora de Valores e Câmbio S.A.'),
(	128	,'MS BANK S.A. BCO DE CÂMBIO','MS Bank S.A. Banco de Câmbio'),
(	129	,'UBS BRASIL BI S.A.','UBS Brasil Banco de Investimento S.A.'),
(	130	,'CARUANA SCFI','CARUANA S.A. - SOCIEDADE DE CRÉDITO, FINANCIAMENTO E INVESTIMENTO'),
(	131	,'TULLETT PREBON BRASIL CVC LTDA','TULLETT PREBON BRASIL CORRETORA DE VALORES E CÂMBIO LTDA'),
(	132	,'ICBC DO BRASIL BM S.A.','ICBC do Brasil Banco Múltiplo S.A.'),
(	133	,'CRESOL CONFEDERAÇÃO','CONFEDERAÇÃO NACIONAL DAS COOPERATIVAS CENTRAIS DE CRÉDITO E ECONOMIA FAMILIAR E'),
(	134	,'BGC LIQUIDEZ DTVM LTDA','BGC LIQUIDEZ DISTRIBUIDORA DE TÍTULOS E VALORES MOBILIÁRIOS LTDA'),
(	136	,'UNICRED','CONFEDERAÇÃO NACIONAL DAS COOPERATIVAS CENTRAIS UNICRED LTDA. - UNICRED DO BRASI'),
(	138	,'GET MONEY CC LTDA','Get Money Corretora de Câmbio S.A.'),
(	139	,'INTESA SANPAOLO BRASIL S.A. BM','Intesa Sanpaolo Brasil S.A. - Banco Múltiplo'),
(	140	,'EASYNVEST - TÍTULO CV SA','Easynvest - Título Corretora de Valores SA'),
(	142	,'BROKER BRASIL CC LTDA.','Broker Brasil Corretora de Câmbio Ltda.'),
(	143	,'TREVISO CC S.A.','Treviso Corretora de Câmbio S.A.'),
(	144	,'BEXS BCO DE CAMBIO S.A.','BEXS BANCO DE CÂMBIO S/A'),
(	145	,'LEVYCAM CCV LTDA','LEVYCAM - CORRETORA DE CAMBIO E VALORES LTDA.'),
(	146	,'GUITTA CC LTDA','GUITTA CORRETORA DE CAMBIO LTDA.'),
(	149	,'FACTA S.A. CFI','Facta Financeira S.A. - Crédito Financiamento e Investimento'),
(	157	,'ICAP DO BRASIL CTVM LTDA.','ICAP do Brasil Corretora de Títulos e Valores Mobiliários Ltda.'),
(	159	,'CASA CREDITO S.A. SCM','Casa do Crédito S.A. Sociedade de Crédito ao Microempreendedor'),
(	163	,'COMMERZBANK BRASIL S.A. - BCO MÚLTIPLO','Commerzbank Brasil S.A. - Banco Múltiplo'),
(	169	,'BCO OLÉ BONSUCESSO CONSIGNADO S.A.','Banco Olé Bonsucesso Consignado S.A.'),
(	173	,'BRL TRUST DTVM SA','BRL Trust Distribuidora de Títulos e Valores Mobiliários S.A.'),
(	174	,'PERNAMBUCANAS FINANC S.A. CFI','PERNAMBUCANAS FINANCIADORA S.A. - CRÉDITO, FINANCIAMENTO E INVESTIMENTO'),
(	177	,'GUIDE','Guide Investimentos S.A. Corretora de Valores'),
(	180	,'CM CAPITAL MARKETS CCTVM LTDA','CM CAPITAL MARKETS CORRETORA DE CÂMBIO, TÍTULOS E VALORES MOBILIÁRIOS LTDA'),
(	183	,'SOCRED SA - SCMEPP','SOCRED S.A. - SOCIEDADE DE CRÉDITO AO MICROEMPREENDEDOR E À EMPRESA DE PEQUENO P'),
(	184	,'BCO ITAÚ BBA S.A.','Banco Itaú BBA S.A.'),
(	188	,'ATIVA S.A. INVESTIMENTOS CCTVM','ATIVA INVESTIMENTOS S.A. CORRETORA DE TÍTULOS, CÂMBIO E VALORES'),
(	189	,'HS FINANCEIRA','HS FINANCEIRA S/A CREDITO, FINANCIAMENTO E INVESTIMENTOS'),
(	190	,'SERVICOOP','SERVICOOP - COOPERATIVA DE CRÉDITO DOS SERVIDORES PÚBLICOS ESTADUAIS DO RIO GRAN'),
(	191	,'NOVA FUTURA CTVM LTDA.','Nova Futura Corretora de Títulos e Valores Mobiliários Ltda.'),
(	194	,'PARMETAL DTVM LTDA','PARMETAL DISTRIBUIDORA DE TÍTULOS E VALORES MOBILIÁRIOS LTDA'),
(	196	,'FAIR CC S.A.','FAIR CORRETORA DE CAMBIO S.A.'),
(	197	,'STONE PAGAMENTOS S.A.','Stone Pagamentos S.A.'),
(	208	,'BANCO BTG PACTUAL S.A.','Banco BTG Pactual S.A.'),
(	212	,'BANCO ORIGINAL','Banco Original S.A.'),
(	213	,'BCO ARBI S.A.','Banco Arbi S.A.'),
(	217	,'BANCO JOHN DEERE S.A.','Banco John Deere S.A.'),
(	218	,'BCO BS2 S.A.','Banco BS2 S.A.'),
(	222	,'BCO CRÉDIT AGRICOLE BR S.A.','BANCO CRÉDIT AGRICOLE BRASIL S.A.'),
(	224	,'BCO FIBRA S.A.','Banco Fibra S.A.'),
(	233	,'BANCO CIFRA','Banco Cifra S.A.'),
(	237	,'BCO BRADESCO S.A.','Banco Bradesco S.A.'),
(	241	,'BCO CLASSICO S.A.','BANCO CLASSICO S.A.'),
(	243	,'BCO MÁXIMA S.A.','Banco Máxima S.A.'),
(	246	,'BCO ABC BRASIL S.A.','Banco ABC Brasil S.A.'),
(	249	,'BANCO INVESTCRED UNIBANCO S.A.','Banco Investcred Unibanco S.A.'),
(	250	,'BCV','BCV - BANCO DE CRÉDITO E VAREJO S.A.'),
(	253	,'BEXS CC S.A.','Bexs Corretora de Câmbio S/A'),
(	254	,'PARANA BCO S.A.','PARANÁ BANCO S.A.'),
(	259	,'MONEYCORP BCO DE CÂMBIO S.A.','MONEYCORP BANCO DE CÂMBIO S.A.'),
(	260	,'NU PAGAMENTOS S.A.','Nu Pagamentos S.A.'),
(	265	,'BCO FATOR S.A.','Banco Fator S.A.'),
(	266	,'BCO CEDULA S.A.','BANCO CEDULA S.A.'),
(	268	,'BARI CIA HIPOTECÁRIA','BARI COMPANHIA HIPOTECÁRIA'),
(	269	,'BCO HSBC S.A.','BANCO HSBC S.A.'),
(	270	,'SAGITUR CC LTDA','Sagitur Corretora de Câmbio Ltda.'),
(	271	,'IB CCTVM S.A.','IB Corretora de Câmbio, Títulos e Valores Mobiliários S.A.'),
(	272	,'AGK CC S.A.','AGK CORRETORA DE CAMBIO S.A.'),
(	273	,'CCR DE SÃO MIGUEL DO OESTE','Cooperativa de Crédito Rural de São Miguel do Oeste - Sulcredi/São Miguel'),
(	274	,'MONEY PLUS SCMEPP LTDA','MONEY PLUS SOCIEDADE DE CRÉDITO AO MICROEMPREENDEDOR E A EMPRESA DE PEQUENO PORT'),
(	279	,'CCR DE PRIMAVERA DO LESTE','COOPERATIVA DE CREDITO RURAL DE PRIMAVERA DO LESTE'),
(	280	,'AVISTA S.A. CFI','Avista S.A. Crédito, Financiamento e Investimento'),
(	281	,'CCR COOPAVEL','Cooperativa de Crédito Rural Coopavel'),
(	283	,'RB CAPITAL INVESTIMENTOS DTVM LTDA.','RB CAPITAL INVESTIMENTOS DISTRIBUIDORA DE TÍTULOS E VALORES MOBILIÁRIOS LIMITADA'),
(	285	,'FRENTE CC LTDA.','Frente Corretora de Câmbio Ltda.'),
(	286	,'CCR DE OURO','COOPERATIVA DE CRÉDITO RURAL DE OURO SULCREDI/OURO'),
(	288	,'CAROL DTVM LTDA.','CAROL DISTRIBUIDORA DE TITULOS E VALORES MOBILIARIOS LTDA.'),
(	289	,'DECYSEO CC LTDA.','DECYSEO CORRETORA DE CAMBIO LTDA.'),
(	290	,'PAGSEGURO','Pagseguro Internet S.A.'),
(	292	,'BS2 DTVM S.A.','BS2 Distribuidora de Títulos e Valores Mobiliários S.A.'),
(	293	,'LASTRO RDV DTVM LTDA','Lastro RDV Distribuidora de Títulos e Valores Mobiliários Ltda.'),
(	296	,'VISION S.A. CC','VISION S.A. CORRETORA DE CAMBIO'),
(	298	,'VIPS CC LTDA.','Vips Corretora de Câmbio Ltda.'),
(	299	,'SOROCRED CFI S.A.','SOROCRED CRÉDITO, FINANCIAMENTO E INVESTIMENTO S.A.'),
(	300	,'BCO LA NACION ARGENTINA','Banco de la Nacion Argentina'),
(	301	,'BPP IP S.A.','BPP Instituição de Pagamento S.A.'),
(	306	,'PORTOPAR DTVM LTDA','PORTOPAR DISTRIBUIDORA DE TITULOS E VALORES MOBILIARIOS LTDA.'),
(	307	,'TERRA INVESTIMENTOS DTVM','Terra Investimentos Distribuidora de Títulos e Valores Mobiliários Ltda.'),
(	309	,'CAMBIONET CC LTDA','CAMBIONET CORRETORA DE CÂMBIO LTDA.'),
(	310	,'VORTX DTVM LTDA.','VORTX DISTRIBUIDORA DE TITULOS E VALORES MOBILIARIOS LTDA.'),
(	315	,'PI DTVM S.A.','PI Distribuidora de Títulos e Valores Mobiliários S.A.'),
(	318	,'BCO BMG S.A.','Banco BMG S.A.'),
(	319	,'OM DTVM LTDA','OM DISTRIBUIDORA DE TÍTULOS E VALORES MOBILIÁRIOS LTDA'),
(	320	,'BCO CCB BRASIL S.A.','China Construction Bank (Brasil) Banco Múltiplo S/A'),
(	321	,'CREFAZ SCMEPP LTDA','CREFAZ SOCIEDADE DE CRÉDITO AO MICROEMPREENDEDOR E A EMPRESA DE PEQUENO PORTE LT'),
(	322	,'CCR DE ABELARDO LUZ','Cooperativa de Crédito Rural de Abelardo Luz - Sulcredi/Crediluz'),
(	323	,'MERCADO PAGO','MERCADOPAGO.COM REPRESENTACOES LTDA.'),
(	325	,'ÓRAMA DTVM S.A.','Órama Distribuidora de Títulos e Valores Mobiliários S.A.'),
(	326	,'PARATI - CFI S.A.','PARATI - CREDITO, FINANCIAMENTO E INVESTIMENTO S.A.'),
(	329	,'QI SCD S.A.','QI Sociedade de Crédito Direto S.A.'),
(	330	,'BANCO BARI S.A.','BANCO BARI DE INVESTIMENTOS E FINANCIAMENTOS S.A.'),
(	331	,'FRAM CAPITAL DTVM S.A.','Fram Capital Distribuidora de Títulos e Valores Mobiliários S.A.'),
(	332	,'ACESSO','Acesso Soluções de Pagamento S.A.'),
(	335	,'BANCO DIGIO','Banco Digio S.A.'),
(	336	,'BCO C6 S.A.','Banco C6 S.A.'),
(	340	,'SUPER PAGAMENTOS E ADMINISTRACAO DE MEIOS ELETRONICOS S.A.','Super Pagamentos e Administração de Meios Eletrônicos S.A.'),
(	341	,'ITAÚ UNIBANCO S.A.','ITAÚ UNIBANCO S.A.'),
(	342	,'CREDITAS SCD','Creditas Sociedade de Crédito Direto S.A.'),
(	343	,'FFA SCMEPP LTDA.','FFA SOCIEDADE DE CRÉDITO AO MICROEMPREENDEDOR E À EMPRESA DE PEQUENO PORTE LTDA.'),
(	348	,'BCO XP S.A.','Banco XP S.A.'),
(	349	,'AMAGGI S.A. CFI','AMAGGI S.A. - CRÉDITO, FINANCIAMENTO E INVESTIMENTO'),
(	352	,'TORO CTVM LTDA','TORO CORRETORA DE TÍTULOS E VALORES MOBILIÁRIOS LTDA'),
(	354	,'NECTON INVESTIMENTOS S.A CVM','NECTON INVESTIMENTOS S.A. CORRETORA DE VALORES MOBILIÁRIOS E COMMODITIES'),
(	355	,'ÓTIMO SCD S.A.','ÓTIMO SOCIEDADE DE CRÉDITO DIRETO S.A.'),
(	364	,'GERENCIANET PAGTOS BRASIL LTDA','GERENCIANET PAGAMENTOS DO BRASIL LTDA'),
(	365	,'SOLIDUS S.A. CCVM','SOLIDUS S.A. CORRETORA DE CAMBIO E VALORES MOBILIARIOS'),
(	367	,'VITREO DTVM S.A.','VITREO DISTRIBUIDORA DE TÍTULOS E VALORES MOBILIÁRIOS S.A.'),
(	370	,'BCO MIZUHO S.A.','Banco Mizuho do Brasil S.A.'),
(	371	,'WARREN CVMC LTDA','WARREN CORRETORA DE VALORES MOBILIÁRIOS E CÂMBIO LTDA.'),
(	373	,'UP.P SEP S.A.','UP.P SOCIEDADE DE EMPRÉSTIMO ENTRE PESSOAS S.A.'),
(	376	,'BCO J.P. MORGAN S.A.','BANCO J.P. MORGAN S.A.'),
(	389	,'BCO MERCANTIL DO BRASIL S.A.','Banco Mercantil do Brasil S.A.'),
(	394	,'BCO BRADESCO FINANC. S.A.','Banco Bradesco Financiamentos S.A.'),
(	399	,'KIRTON BANK','Kirton Bank S.A. - Banco Múltiplo'),
(	412	,'BCO CAPITAL S.A.','BANCO CAPITAL S.A.'),
(	422	,'BCO SAFRA S.A.','Banco Safra S.A.'),
(	456	,'BCO MUFG BRASIL S.A.','Banco MUFG Brasil S.A.'),
(	464	,'BCO SUMITOMO MITSUI BRASIL S.A.','Banco Sumitomo Mitsui Brasileiro S.A.'),
(	473	,'BCO CAIXA GERAL BRASIL S.A.','Banco Caixa Geral - Brasil S.A.'),
(	477	,'CITIBANK N.A.','Citibank N.A.'),
(	479	,'BCO ITAUBANK S.A.','Banco ItauBank S.A.'),
(	487	,'DEUTSCHE BANK S.A.BCO ALEMAO','DEUTSCHE BANK S.A. - BANCO ALEMAO'),
(	492	,'ING BANK N.V.','ING Bank N.V.'),
(	495	,'BCO LA PROVINCIA B AIRES BCE','Banco de La Provincia de Buenos Aires'),
(	505	,'BCO CREDIT SUISSE S.A.','Banco Credit Suisse (Brasil) S.A.'),
(	545	,'SENSO CCVM S.A.','SENSO CORRETORA DE CAMBIO E VALORES MOBILIARIOS S.A'),
(	600	,'BCO LUSO BRASILEIRO S.A.','Banco Luso Brasileiro S.A.'),
(	604	,'BCO INDUSTRIAL DO BRASIL S.A.','Banco Industrial do Brasil S.A.'),
(	610	,'BCO VR S.A.','Banco VR S.A.'),
(	611	,'BCO PAULISTA S.A.','Banco Paulista S.A.'),
(	612	,'BCO GUANABARA S.A.','Banco Guanabara S.A.'),
(	613	,'OMNI BANCO S.A.','Omni Banco S.A.'),
(	623	,'BANCO PAN','Banco Pan S.A.'),
(	626	,'BCO FICSA S.A.','BANCO FICSA S.A.'),
(	630	,'SMARTBANK','Banco Smartbank S.A.'),
(	633	,'BCO RENDIMENTO S.A.','Banco Rendimento S.A.'),
(	634	,'BCO TRIANGULO S.A.','BANCO TRIANGULO S.A.'),
(	637	,'BCO SOFISA S.A.','BANCO SOFISA S.A.'),
(	643	,'BCO PINE S.A.','Banco Pine S.A.'),
(	652	,'ITAÚ UNIBANCO HOLDING S.A.','Itaú Unibanco Holding S.A.'),
(	653	,'BCO INDUSVAL S.A.','BANCO INDUSVAL S.A.'),
(	654	,'BCO A.J. RENNER S.A.','BANCO A.J. RENNER S.A.'),
(	655	,'BCO VOTORANTIM S.A.','Banco Votorantim S.A.'),
(	707	,'BCO DAYCOVAL S.A','Banco Daycoval S.A.'),
(	712	,'BCO OURINVEST S.A.','Banco Ourinvest S.A.'),
(	739	,'BCO CETELEM S.A.','Banco Cetelem S.A.'),
(	741	,'BCO RIBEIRAO PRETO S.A.','BANCO RIBEIRAO PRETO S.A.'),
(	743	,'BANCO SEMEAR','Banco Semear S.A.'),
(	745	,'BCO CITIBANK S.A.','Banco Citibank S.A.'),
(	746	,'BCO MODAL S.A.','Banco Modal S.A.'),
(	747	,'BCO RABOBANK INTL BRASIL S.A.','Banco Rabobank International Brasil S.A.'),
(	748	,'BCO COOPERATIVO SICREDI S.A.','BANCO COOPERATIVO SICREDI S.A.'),
(	751	,'SCOTIABANK BRASIL','Scotiabank Brasil S.A. Banco Múltiplo'),
(	752	,'BCO BNP PARIBAS BRASIL S A','Banco BNP Paribas Brasil S.A.'),
(	753	,'NOVO BCO CONTINENTAL S.A. - BM','Novo Banco Continental S.A. - Banco Múltiplo'),
(	754	,'BANCO SISTEMA','Banco Sistema S.A.'),
(	755	,'BOFA MERRILL LYNCH BM S.A.','Bank of America Merrill Lynch Banco Múltiplo S.A.'),
(	756	,'BANCOOB','BANCO COOPERATIVO DO BRASIL S.A. - BANCOOB'),
(	757	,'BCO KEB HANA DO BRASIL S.A.','BANCO KEB HANA DO BRASIL S.A.')
;
