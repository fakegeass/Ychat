# -*- coding: utf-8 -*- 

###########################################################################
## Python code generated with wxFormBuilder (version Jul 18 2018)
## http://www.wxformbuilder.org/
##
## PLEASE DO *NOT* EDIT THIS FILE!
###########################################################################

import wx
import wx.xrc

from socket import *
HOST ='213.213.0.9'
PORT = 8000
BUFFSIZE = 2048
ADDR = (HOST,PORT)

from Chat import *

###########################################################################
## Class Login
###########################################################################

class Login ( wx.Frame ):
	
	def __init__( self, parent ):
		wx.Frame.__init__ ( self, parent, id = wx.ID_ANY, title = u"登录到聊天室", pos = wx.DefaultPosition, size = wx.Size( 400,200 ), style = wx.DEFAULT_FRAME_STYLE|wx.TAB_TRAVERSAL )
		
		self.SetSizeHints( wx.DefaultSize, wx.DefaultSize )
		
		bSizer1 = wx.BoxSizer( wx.VERTICAL )
		
		self.ServerName = wx.StaticText( self, wx.ID_ANY, u"欢迎来到ACGN精英论坛，死宅出品，必属精品！", wx.DefaultPosition, wx.DefaultSize, 0 )
		self.ServerName.Wrap( -1 )
		
		bSizer1.Add( self.ServerName, 0, wx.ALL, 5 )
		
		gSizer1 = wx.GridSizer( 0, 2, 0, 0 )
		
		gSizer1.SetMinSize( wx.Size( -1,100 ) ) 
		self.staticUid = wx.StaticText( self, wx.ID_ANY, u"用户名：", wx.DefaultPosition, wx.DefaultSize, 0 )
		self.staticUid.Wrap( -1 )
		
		self.staticUid.SetMinSize( wx.Size( 100,-1 ) )
		
		gSizer1.Add( self.staticUid, 0, wx.ALL, 5 )
		
		self.Uid = wx.TextCtrl( self, wx.ID_ANY, wx.EmptyString, wx.DefaultPosition, wx.DefaultSize, 0 )
		gSizer1.Add( self.Uid, 0, wx.ALL, 5 )
		
		self.StaticPs = wx.StaticText( self, wx.ID_ANY, u"密码：", wx.DefaultPosition, wx.DefaultSize, 0 )
		self.StaticPs.Wrap( -1 )
		
		gSizer1.Add( self.StaticPs, 0, wx.ALL, 5 )
		
		self.Ps = wx.TextCtrl( self, wx.ID_ANY, wx.EmptyString, wx.DefaultPosition, wx.DefaultSize, 0 )
		gSizer1.Add( self.Ps, 0, wx.ALL, 5 )
		
		
		bSizer1.Add( gSizer1, 1, wx.EXPAND, 5 )
		
		gSizer2 = wx.GridSizer( 0, 2, 0, 0 )
		
		self.Confirm = wx.Button( self, wx.ID_ANY, u"登录", wx.DefaultPosition, wx.DefaultSize, 0 )
		gSizer2.Add( self.Confirm, 0, wx.ALL|wx.ALIGN_CENTER_VERTICAL|wx.ALIGN_CENTER_HORIZONTAL, 5 )
		
		self.Cancel = wx.Button( self, wx.ID_ANY, u"取消", wx.DefaultPosition, wx.DefaultSize, 0 )
		gSizer2.Add( self.Cancel, 0, wx.ALL|wx.ALIGN_CENTER_HORIZONTAL, 5 )
		
		
		bSizer1.Add( gSizer2, 1, wx.EXPAND, 5 )
		
		
		self.SetSizer( bSizer1 )
		self.Layout()
		
		self.Centre( wx.BOTH )
		
		# Connect Events
		self.Uid.Bind( wx.EVT_TEXT, self.eve_Uid )
		self.Ps.Bind( wx.EVT_TEXT, self.eve_Ps )
		self.Confirm.Bind( wx.EVT_BUTTON, self.eve_Login )
		self.Cancel.Bind( wx.EVT_BUTTON, self.eve_Exit )
	
	def __del__( self ):
		pass
	
	# Virtual event handlers, overide them in your derived class
	def eve_Uid( self, event ):
		self.id = self.Uid.GetLineText(0)
	
	def eve_Ps( self, event ):
		self.ps = self.Ps.GetLineText(0)
	
	def eve_Login( self, event ):
		self.tctimeClient = socket(AF_INET,SOCK_STREAM)
		self.tctimeClient.connect(ADDR)

		data = '0'+self.id+'\n'
		self.tctimeClient.send(data.encode())
		data = self.tctimeClient.recv(BUFFSIZE).decode().strip()
		print(data)

		if data != '0':
			eve_Exit()

		data = '1'+self.ps+'\n'
		self.tctimeClient.send(data.encode())
		data = self.tctimeClient.recv(BUFFSIZE).decode().strip()

		if data == '0':
			print('login succeed!')

		self.Destroy()
		chat = Chat(parent=None, tctimeClient=self.tctimeClient)
		chat.Show()


			
	def eve_Exit( self, event ):
		self.Destroy()
	

