#! /usr/bin/env python 
import sys

from tinytag import TinyTag

tag = TinyTag.get(sys.argv[1])

if sys.argv[2] == "json":
  print('{"bitrate":%f, "duration":%s}' % (tag.bitrate, tag.duration))
else:
  print('Bitrate:%f|Duration:%s' % (tag.bitrate, tag.duration))
