Like [outset](https://github.com/chilcote/outset) for [yo](https://github.com/sheagcraig/yo)

yo-yo traverses a directory for json or plist formatted code parameters for Yo, and creates notifications using yo. 
Installing a LaunchAgent to watch a directory and trigger yo-yo allows applications running as root to deliver notifications to the session.
yo-yo deletes the notification file after successfuly executing `yo`, promising at most once delivery. This won't work in a multi user context.

Notification file example:
The keys in the json or plist file mirror yo's cli flags.

```plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>title</key>
    <string>Foo</string>
    <key>subtitle</key>
    <string>bar</string>
</dict>
</plist>
```
