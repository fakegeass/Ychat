from Login import *
from Chat import *
import wx
import wx.xrc
import time


app = wx.App(False)

frame = Login(None)
frame.Show()
app.MainLoop()