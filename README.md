# tstomkvall

A small utility program that automatically transcodes all .ts files in a directory to .mkv files losslessly. It uses [FFmpeg](https://ffmpeg.org/) to copy all streams (i.e. video, audio, subtitles, etc.) from the source file to a Matroska (.mkv) container and does not touch anything in the process. Mainly used for batch transcoding downloaded internet streams stored in .ts format.

Requires [FFmpeg](https://ffmpeg.org/) in PATH.
