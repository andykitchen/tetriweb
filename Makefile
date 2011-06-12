include $(GOROOT)/src/Make.inc
 
TARG=tetriweb
GOFILES=main.go\
        shapes.go \
        board.go \
        gameserver.go

include $(GOROOT)/src/Make.cmd
