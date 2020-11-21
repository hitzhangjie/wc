README
==============================================================================

wc is an totally alternative to Unix/Linux wc utility.

wc takes following factors into consideration:
- Non-English words, like Chinese;
- Count by words, bytes, lines.

It supports all the options in Unix/Linux wc, and adds `--sort` options.

```bash
find -iname "*zh_CN.md" -print0 | grep -z -v _book | grep -z -v node_modules |  wc -c -m -w -l --sort lines --files0-from - 
0	0	0	0	./8-develop-sym-debugger/2-debug-line.zh_CN.md			
0	0	0	0	./9-others.zh_CN.md						
3	9	95	95	./10-contributors.zh_CN.md					
5	3	135	397	./6-debugger-skeleton/README.zh_CN.md				
6	8	146	354	./5-dwarf/4-other-shrink-data.zh_CN.md				
7	12	112	270	./4-basics/README.zh_CN.md					
8	8	191	485	./5-dwarf/4-other-macro-info.zh_CN.md				
8	10	258	554	./5-dwarf/4-other-varlen-data.zh_CN.md				
8	18	450	976	./5-dwarf/README.zh_CN.md					
9	8	418	1134	./2-preface.zh_CN.md						
11	9	387	897	./5-dwarf/3-die-readme.zh_CN.md					
14	58	1085	2127	./3-terms.zh_CN.md						
17	30	521	1025	./5-dwarf/5-summary.zh_CN.md					
17	50	344	348	./7-develop-inst-debugger/20-memory-registers.zh_CN.md		
18	31	1050	2620	./5-dwarf/4-other-accelerated-access.zh_CN.md			
19	34	825	1833	./5-dwarf/2-structure.zh_CN.md					
19	21	616	1416	./5-dwarf/3-die-encoding.zh_CN.md				
25	29	700	1888	./4-basics/1-purposes.zh_CN.md					
32	53	1120	2330	./5-dwarf/4-other-elf-sections.zh_CN.md				
41	82	1792	3806	./5-dwarf/1-history.zh_CN.md					
47	54	1655	3789	./5-dwarf/4-other-readme.zh_CN.md				
49	109	1811	4249	./1-introduction.zh_CN.md					
56	174	2934	3362	./SUMMARY.zh_CN.md						
59	100	1755	4193	./11-contributing.zh_CN.md					
64	82	2113	4415	./5-dwarf/3-die-desc-code.zh_CN.md				
75	137	2504	4480	./5-dwarf/3-die-intro.zh_CN.md					
89	140	2096	4460	./4-basics/3-countertactics.zh_CN.md				
93	139	1428	2002	./7-develop-inst-debugger/8-clearall.zh_CN.md			
114	504	3336	3342	./8-develop-sym-debugger/1-symtable & linetable.zh_CN.md	
124	199	2080	3090	./7-develop-inst-debugger/7-clear.zh_CN.md			
126	226	2789	5165	./7-develop-inst-debugger/1-process_start.zh_CN.md		
143	704	6284	7962	./7-develop-inst-debugger/README.zh_CN.md			
147	310	3570	5414	./7-develop-inst-debugger/10-continue.zh_CN.md			
148	253	3342	5840	./7-develop-inst-debugger/5-breakpoint.zh_CN.md			
163	265	6120	12532	./6-debugger-skeleton/1-debugger_skeleton.zh_CN.md		
164	1422	8966	9002	./8-develop-sym-debugger/3-call-frame.zh_CN.md			
170	406	4250	6281	./7-develop-inst-debugger/9-step.zh_CN.md			
190	490	7640	14962	./5-dwarf/3-die-desc-datatype.zh_CN.md				
195	275	7502	16460	./4-basics/2-dependencies.zh_CN.md				
199	362	4007	6077	./7-develop-inst-debugger/6-breakpoints.zh_CN.md		
200	354	6861	13797	./5-dwarf/4-other-lineno-table.zh_CN.md				
249	806	6944	9596	./7-develop-inst-debugger/2-process_attach.zh_CN.md		
285	543	7039	10745	./6-debugger-skeleton/2-debugger_demo.zh_CN.md			
305	633	7495	11145	./7-develop-inst-debugger/3-process_start_attach.zh_CN.md	
305	634	7124	11646	./7-develop-inst-debugger/4-disassemble.zh_CN.md		
419	857	17106	33678	./5-dwarf/4-other-callframe-info.zh_CN.md			
507	1881	13658	13660	./8-develop-sym-debugger/x-gopkg.zh_CN.md			
4952	12532	152654	253899	total
```