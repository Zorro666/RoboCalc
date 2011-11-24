# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include ${GOROOT}/src/Make.inc

C_CPP_COMMON_COMPILE_FLAGS:= -g -Wall -Wextra -Wuninitialized -Winit-self -Wstrict-aliasing -Wfloat-equal -Wshadow -Wconversion -Werror -fpack-struct=4

C_COMPILE:=gcc -c
C_COMPILE_FLAGS:=-ansi -pedantic-errors

C_LINK:=g++
C_LINK_FLAGS:=-g -lm -lpthread

GO_PROJECTS:=robocalc

C_PROJECTS:=robocalcC

all: $(GO_PROJECTS) $(C_PROJECTS)

define upperString
$(shell echo $1 | tr [a-z] [A-Z] )
endef

define GO_PROJECT_template
$2_SRCFILES += $1.go
$2_SRCFILES += $($2_DEPENDS)
$2_DEPEND_OBJS:=$($2_DEPENDS:.go=.8)

$2_OBJFILE:=$1.8
$2_OBJFILES:=$$($2_SRCFILES:.go=.8)
$2_FMTFILES:=$$($2_SRCFILES:.go=.fmt.tmp)

GO_SRCFILES += $$($2_SRCFILES)
GO_OBJFILES += $$($2_OBJFILES)
GO_FMTFILES += $$($2_FMTFILES)

GO_TARGETS += $1

$$($2_OBJFILE): $$($2_DEPEND_OBJS) $1.go
$1: $$($2_OBJFILES) 
endef

define C_PROJECT_template
$2_SRCFILES += $1.c
$2_SRCFILES += $($2_DEPENDS)
$2_DEPEND_OBJS:=$($2_DEPENDS:.c=.o)

$2_OBJFILE:=$1.o
$2_OBJFILES:=$$($2_SRCFILES:.c=.o)

C_SRCFILES += $$($2_SRCFILES)
C_OBJFILES += $$($2_OBJFILES)

C_TARGETS += $1

$$($2_OBJFILE): $$($2_DEPEND_OBJS) $1.c
$1: $$($2_OBJFILES) 
endef
     
$(foreach project,$(GO_PROJECTS),$(eval $(call GO_PROJECT_template,$(project),$(call upperString,$(project)))))

$(foreach project,$(C_PROJECTS),$(eval $(call C_PROJECT_template,$(project),$(call upperString,$(project)))))

test:
	@echo GO_PROJECTS=$(GO_PROJECTS)
	@echo GO_TARGETS=$(GO_TARGETS)
	@echo GO_SRCFILES=$(GO_SRCFILES)
	@echo GO_OBJFILES=$(GO_OBJFILES)
	@echo GO_FMTFILES=$(GO_FMTFILES)
	@echo C_PROJECTS=$(C_PROJECTS)
	@echo C_TARGETS=$(C_TARGETS)
	@echo C_SRCFILES=$(C_SRCFILES)
	@echo C_OBJFILES=$(C_OBJFILES)

%.8: %.go
	@echo Go Compiling $<
	@$(GC) $<

%: %.8
	@echo Go Linking $<
	@$(LD) -o $@ $<

%.o: %.c
	@echo C Compiling $<
	@$(C_COMPILE) -MMD $(C_CPP_COMMON_COMPILE_FLAGS) $(C_COMPILE_FLAGS) -o $*.o $<

%: %.o
	@echo C Linking $<
	@$(C_LINK) -o $@ $^ $(C_LINK_FLAGS)

.PHONY: all clean nuke format
.SUFFIXES:            # Delete the default suffixes

FORCE:

clean: FORCE
	rm -f $(GO_OBJFILES)
	rm -f $(C_OBJFILES)

nuke: clean
	rm -f $(GO_TARGETS)
	rm -f $(C_TARGETS)

%.fmt.tmp: %.go
	gofmt -tabwidth=4 -w=true $<
	@rm -f $@

format: FORCE $(FMTFILES)
