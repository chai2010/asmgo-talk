global main
extern take_many

bits 64

section .text
main:
	; take_many(1, 2, 3, 4, 5, 6, 7, 8);
	push 8
	push 7
	mov r9, 6
	mov r8, 5
	mov rcx, 4
	mov rdx, 3
	mov rsi, 2
	mov rdi, 1
	call take_many
	add rsp, 16

	; return 0;
	mov rax, 0
	ret
