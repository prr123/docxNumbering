# docx read

A MS word docx file is a zipped file of several files.  

Structure:  

'[Content_Types].xml'   _rels   docProps   word
                           |       |         |
						.rels      |         |
                           app.xml  core.xml |
document.xml  numbering.xml  settings.xml  fontTable.xml  styles.xml  webSettings.xml  
_rels theme  

## numDecode
program that reads from the unzipped directory tree the file "numbering.xml"   


## numDecode
program that reads from the zipped docx file the file "numbering.xml" .  
It uses the goDocx unpack routine to unzip the docx file.   

