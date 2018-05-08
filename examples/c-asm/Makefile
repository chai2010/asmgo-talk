OBJECTS = func.o main.o
CC = gcc
CFLAGS = -std=c11 -m64 -Wall -Wextra -Werror -c
AS = nasm
#ASFLAGS = -f elf64
ASFLAGS=-f macho64

all: $(OBJECTS)
	gcc -m64 $(OBJECTS) -o main

run: all
	./main

%.o: %.c
	$(CC) $(CFLAGS)  $< -o $@

%.o: %.s
	$(AS) $(ASFLAGS) $< -o $@

clean:
	rm -rf $(OBJECTS) main
