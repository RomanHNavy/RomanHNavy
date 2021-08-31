Python 3.9.1 (default, Dec 11 2020, 09:29:25) [MSC v.1916 64 bit (AMD64)] on win32
Type "help", "copyright", "credits" or "license()" for more information.
>>> import os
>>> import pandas as pd
>>> import openpyxl
>>> import xlsxwriter
>>> from openpyxl import load_workbook
>>> 
>>> #This code will add a DataFrame to an already created excel DataFrame, onto a new MS excel file.
>>> 
>>> #adding invoices from 16_07_2021, to a list of invoices from 15_07_2021
>>> 
>>> wb = openpyxl.load_workbook('16_07_Invoices.xlsx')
>>> wb
<openpyxl.workbook.workbook.Workbook object at 0x0000020DA0775790>
>>> wb.sheetnames
['Sheet1']
>>> sheet = wb['Sheet1']
>>> 
>>> Invoice_List = []
>>> Notes_List = []
>>> Status_List = []
>>> 
>>> for i in range(3, 100, 1):
	if(sheet.cell(row=i, column=3).value) == 'Incomplete':
		Invoice_List.append(sheet.cell(row=i, column=1).value)
		Notes_List.append(sheet.cell(row=i, column=2).value)
		Status_List.append(sheet.cell(row=i, column=3).value)

		
>>> Invoice_List
[516843, 6883726]
>>> 
>>> df = pd.DataFrame({'Invoice': Invoice_List, 'Notes': Notes_List, 'Status': Status_List})
>>> 
>>> writer = pd.ExcelWriter('Incomplete_Invoices_Final!.xlsx', engine 'openpyxl')
SyntaxError: invalid syntax
>>> writer = pd.ExcelWriter('Incomplete_Invoices_Final!.xlsx', engine='openpyxl')
>>> writer.book = load_workbook('Incomplete_Invoices_Final.xlsx')
>>> writer.sheets = dict((ws.title, ws) for ws in writer.book.worksheets)
>>> reader = pd.read_excel(r'Incomplete_Invoices_Final.xlsx')
>>> df.to_excel(writer, index=False, header=False, startrow=len(reader)+1)
>>> 
>>> writer.save()
>>> writer.close()
>>> 