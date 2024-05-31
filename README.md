# ASCII-gif-generator
Side project to make ASCII gifs. We are using Golang to do this project.


## References

- Python Gif to ASCII: https://github.com/tpoff/Python-Gif-Ascii-Animator/blob/master/Gif_Ascii_Animator.py
- DUCK: https://github.com/sahilgajjar/ascii_art/blob/main/duck.c 
- ASCIIMATICS: https://github.com/peterbrittain/asciimatics

## Attack Strategy

Simple-to-Complex Roadmap:
- start with one frame moving across screen
- change to like two frames
- create a set of frames for a single gif
- repeat above for two or three example gifs
- *convert gif to frames for any arbitrary gif* -- another project entirely

Runtime Procedure:
- go run gifmaker.go -giftype="name"