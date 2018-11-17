# -*- coding: utf-8 -*- 

import wx
import wx.xrc
import wx.richtext

BUFFSIZE = 2048

class Chat ( wx.Frame ):
	
	def __init__( self, parent, tctimeClient ):
		self.tctimeClient = tctimeClient
		self.tctimeClient.settimeout(0.05)
		self.receiveWord = ''
		self.word = ''
		self.newWord = ''

		wx.Frame.__init__ ( self, parent, id = wx.ID_ANY, title = u"聊天室", pos = wx.DefaultPosition, size = wx.Size( 500,300 ), style = wx.DEFAULT_FRAME_STYLE|wx.TAB_TRAVERSAL )
		
		self.SetSizeHints( wx.DefaultSize, wx.DefaultSize )
		
		bSizer3 = wx.BoxSizer( wx.VERTICAL )
		
		self.ChatFrame = wx.richtext.RichTextCtrl( self, wx.ID_ANY, wx.EmptyString, wx.DefaultPosition, wx.DefaultSize, 0|wx.VSCROLL|wx.HSCROLL|wx.NO_BORDER|wx.WANTS_CHARS|wx.TE_READONLY )
		self.ChatFrame.SetMinSize( wx.Size( -1,180 ) )
		
		bSizer3.Add( self.ChatFrame, 1, wx.EXPAND |wx.ALL, 5 )
		
		fgSizer2 = wx.FlexGridSizer( 0, 2, 0, 0 )
		fgSizer2.SetFlexibleDirection( wx.BOTH )
		fgSizer2.SetNonFlexibleGrowMode( wx.FLEX_GROWMODE_SPECIFIED )
		
		self.MyWords = wx.TextCtrl( self, wx.ID_ANY, wx.EmptyString, wx.DefaultPosition, wx.Size( 350,90 ), 0 )
		fgSizer2.Add( self.MyWords, 0, wx.ALL, 5 )
		
		self.Send = wx.Button( self, wx.ID_ANY, u"发送", wx.DefaultPosition, wx.DefaultSize, 0 )
		fgSizer2.Add( self.Send, 0, wx.ALL|wx.ALIGN_CENTER_HORIZONTAL|wx.ALIGN_CENTER_VERTICAL, 5 )
		
		
		bSizer3.Add( fgSizer2, 1, wx.EXPAND, 5 )
		
		
		self.SetSizer( bSizer3 )
		self.Layout()
		
		self.Centre( wx.BOTH )
		
		# Connect Events
		self.Bind( wx.EVT_IDLE, self.evn_Idle, id=wx.ID_ANY)
		self.MyWords.Bind( wx.EVT_TEXT, self.evt_MyWords )
		self.Send.Bind( wx.EVT_BUTTON, self.evt_Send )
	
	def __del__( self ):
		self.tctimeClient.close()
	
	
	# Virtual event handlers, overide them in your derived class
	def evn_Idle( self, event ):
		try:
			data = self.tctimeClient.recv(BUFFSIZE).decode()
		except:
			data = ''
		if data:
			self.word += data.strip('0')
			self.ChatFrame.SetValue(str(self.word))
			self.Refresh()
	
	def evt_MyWords( self, event ):
		self.newWord = self.MyWords.GetLineText(0)
	
	def evt_Send( self, event ):
		data = '2'+self.newWord+'\n'
		self.tctimeClient.send(data.encode())
		self.MyWords.Clear()