# mindb

[Local network/localhost does not work 

From the sharpWebview readme:

[webview/webview_csharp: C# bindings for zserge/webview - Batteries included](https://github.com/webview/webview_csharp)

The Edge webview uses a UWP application context on windows. UWP applications disallow loopbacks. For development purpose it is necessary to run the following command in an administrative command prompt:

```
CheckNetIsolation.exe LoopbackExempt -a -n="Microsoft.Win32WebViewHost_cw5n1h2txyewy"
```
This adds the Edge Webview Host to the exception list of this limitation. Your best bet for application distribution is to create an installer which executes this command on installation.

Check, if this works for you.


* [Mage :: Mage](https://magefile.org/)

Linux:

```
sudo apt-get install -y webkit2gtk-4.0
```