## Static Compilation,Coff,DWARF,Debug Signals,ELF,symbolic debugging 

Runtime vs compiler time:  Compile-time is the time at which the source code is converted into an executable code while the run time is the time at which the executable code is started running

### Dynamic Linking:
* Dynamic linking is a process used in computer systems where external libraries or modules are linked to a program at runtime rather than at compiler time.
*  In this approach, the executable file of the program does not include the actual code for the libraries or modules it depends on; instead, it contains references to the shared libraries that are loaded into memory when the program is executed.
* Key Features of Dynamic Linking:
    * Efficiency: Reduces the size of executable files since the libraries are not included in the binary itself. Multiple programs can share a single copy of the library in memory, conserving resources.
    * Modularity: Allows for easier updates and maintenance. If a shared library is updated or patched, all programs using that library can benefit from the update without needing to be recompiled.
    * Flexibility: Enables loading different versions of libraries at runtime, which can be useful for backward compatibility or for running different configurations of a program.
    * Often-used libraries (for example the standard system library) need to be stored in only one location, not duplicated in every single binary. 
    * If a library is upgraded or replaced, all programs using it dynamically will immediately benefit from the corrections. 
    * Static builds would have to be re-linked first. The binary executable file size is smaller than its statically linked counterpart.
    * You can update routines in libraries without needing to relink, meaning bug fixes and other changes can be executed without any need to ship a new executable file.
    * There are lower maintenance costs and a reduced need for support.
    * 
* How It Works:
    * Compilation: During compilation, references to external functions or variables are marked as unresolved, and the linker leaves placeholders in the executable.
    * Runtime: When the program is executed, the dynamic linker/loader loads the necessary shared libraries into memory, resolves the references, and links them to the placeholders in the executable.
    * Execution: The program then executes normally, using the functions and data provided by the dynamically linked libraries.
* Disadvantages:
    * Complexity: Increases runtime complexity due to the need to load and link libraries during execution.
    * Dependency Issues: Can lead to "DLL hell" or similar issues where different programs require different versions of the same library, causing conflicts.
    * While it sounds more efficient overall, dynamic linking comes with its own set of risks.
    * The removal of a library can cause the program to break. Or if any changes are made to a library, you might run into compatibility issues and rework the application as a result.
    * If you have an inflexible dynamic linker, you could find the program no longer launches at all or doesn’t work correctly (also known as “DLL hell”).
    * 
Static Linking:
* In a statically built program, no dynamic linking occurs: all the bindings have been done at compile time.
* Static linking is the practice of copying all the libraries your program will need directly into the final executable file. This occurs right at the end of compilation and is performed by a linker.
* First, the linker combines the program code with all required libraries. Once that’s complete, everything is compiled into one executable file. When the program is installed, all libraries required are loaded into memory.
* That’s where the term “static” comes into play. Now that everything is shipped together, you need to perform this whole process again if any changes are made to the external libraries.
* Static builds have a very predictable behavior (because they do not rely on the particular version of libraries available on the final system), and are commonly found in forensic and security tools to avoid possible contamination or malfunction due to broken libraries on the examined machine
* Another benefit of static builds is their portability: once the final executable file has been compiled, it is no longer necessary to keep the library files that the program references, since all the relevant parts are copied into the executable file. As a result, when installing a statically-built program on a computer, the user doesn't have to download and install additional libraries: the program is ready to run.
* It sounds like a long process, but there are several benefits to using static linking:
    * Static linking can simplify the process of distributing binaries to multiple environments or operating systems, as the program already has everything it needs to run.
    * Static linking can result in a slightly faster start-up depending on the program’s complexity. This also results in less chance of compatibility issues.
* But static linking has its downsides, too.
    * Since you’re compiling and shipping all libraries in one executable file, the program can be much larger and use a lot of resources. That includes cache, RAM, and disk space.
    * If any changes occur in external programs, they won’t automatically be reflected in your executable file. You’ll need to perform the whole compilation and linking process again
Otool:
https://stackoverflow.com/questions/55196053/ldd-r-equivalent-on-macos

## COFF:

* The Common Object File Format (COFF) is a specific file format suitable for code debugging.
* The COFF incorporates symbolic procedure, function, variable and constant names information; line number information, breakpoints settings, code highlighter and all the necessary information for effective and fast debugging.
* .cof
* The Common Object File Format (COFF) is a format for executable, object code, and shared library computer files used on Unix systems.
* Object code:
    * Object code is the intermediate, machine-readable output generated by a compiler after processing the source code written in a high-level programming language
    * This code is a low-level representation of the original source code, often in the form of binary or machine code, which the processor can understand directly but is not yet fully executable.
    * Characteristics of Object Code:
        * Machine-Specific: Object code is typically tailored for a specific type of processor or computer architecture.
        * Not Executable: By itself, object code is usually not directly executable because it may lack references to external libraries or system calls needed to run the program. It needs to be linked with other object files or libraries to create a complete executable.
        * Binary Format: Object code is usually in a binary format, making it difficult for humans to read or interpret directly.
    * How Object Code is Produced:
        * Source Code: A programmer writes source code in a high-level language like C, C++, or Java.
        * Compilation: The source code is passed through a compiler, which translates it into object code. The compiler performs syntax analysis, optimization, and code generation during this process.
        * Object File: The compiler outputs one or more object files, typically with a .o, .obj, or .out extension, depending on the system and language.
    * Linkage:
        * Static Linking: Object files can be combined with other object files and libraries using a linker to produce an executable program.
        * Dynamic Linking: Object code may rely on dynamic libraries, which are linked at runtime.
    * For example, if you write a simple C program and compile it, the compiler will generate an object file. This object file contains the machine code corresponding to your C program but is not yet executable. To run the program, the object file must be linked with the necessary runtime libraries to produce the final executable.
* COFF was introduced in Unix System V, replaced the previously used a.out format, and formed the basis for extended specifications such as XCOFF and ECOFF, before being largely replaced by ELF, introduced with SVR4.
* The original Unix object file format a.out is unable to adequately support shared libraries, foreign format identification or explicit address linkage.
* Key Components of a COFF File:
    * Header:
        * The COFF file begins with a header that describes the characteristics of the file.
        * The header includes fields like the number of sections, the timestamp of file creation, and flags indicating the type of file (e.g., relocatable, executable)."Relocatable" refers to a type of object file or code that is not yet fully resolved in terms of memory addresses and can be placed at different locations in memory during the linking or loading phase of program execution.
        * It also specifies the size of the optional header and the number of symbols in the symbol table.
    * Section Table:
        * Immediately following the header, the section table provides information about the various sections of the object file.
        * Each entry in the section table corresponds to a section in the file, such as .text (code), .data (initialized data), and .bss (uninitialized data).
        * Each section has attributes like size, memory location, and file offset where the section's data can be found.
    * Sections:
        * The bulk of the COFF file consists of sections containing the actual data of the program.
        * .text section: Contains the executable code.
        * .data section: Contains initialized global and static variables. A static variable is a variable in computer programming that is allocated "statically" and exists for the entire run of a program
        * .bss section: Represents uninitialized variables; it does not occupy space in the file, as it is initialized to zero at runtime.
        * Other sections may include debug information, relocation entries, and more.
    * Symbol Table:
        * The symbol table lists all symbols (such as variable names, function names, and labels) defined or used in the program.
        * Each symbol entry in the table contains information like the symbol's name, type, section number, value (e.g., address or offset), and storage class (e.g., external, static).
        * The symbol table is used during the linking process to resolve references to functions and variables between different object files.
    * String Table:
        * Names of symbols and sections that exceed a certain length are stored in a separate string table.
        * The symbol table entries then contain offsets pointing to the names in the string table.
        * This allows for more flexible naming without being constrained by fixed-length fields.
    * Relocation Entries:
        * If the object file contains addresses that cannot be fully resolved until linking, it includes relocation entries.
        * ¯Relocation entries indicate where in the object file these addresses are located and how they should be adjusted when the final executable or library is created.
* Advantages of COFF:
    * Portability: COFF was designed to be machine-independent, making it easier to move code between different systems.
    * Modularity: It supports the separation of code into different sections, which aids in modular programming and efficient memory management.
    * Extensibility: COFF can be extended with additional sections to store debugging information, making it a good format for development as well as final deployment.
* A later implementation  of COFF by A&t labs is ELF 
## ELF:
* In computing, the Executable and Linkable Format(ELF, formerly named Extensible Linking Format) is a common standard file format for executable files, object code, shared libraries, and core dumps.
Filename extension	none, .axf, .bin, .elf, .o, .out, .prx, .puff, .ko, .mod, and .so
* By design, the ELF format is flexible, extensible, and cross-platform.
* Types of ELF Files:
* Relocatable File (.o):
    * Generated by the compiler, these files contain code and data that need to be linked with other object files to create an executable or shared library.
* Executable File:
    * A complete, ready-to-run program. It has a valid entry point and is fully linked, with all addresses resolved.
* Shared Object (.so):
    * These are shared libraries that can be loaded and linked at either compile-time or runtime. They allow code to be shared across multiple programs, reducing redundancy.
* Core Dump:
    * A snapshot of a program's memory and state at a particular point in time, typically when the program crashes. Core dumps are used for debugging.
* ELF Structure Overview:
* ELF Header: Describes the file, including type, architecture, and offsets to program and section headers.
* Program Header Table: Describes how to create a process image, used for executables and shared libraries.
* Sections: Various pieces of the program, like code, data, and metadata.
* Section Header Table: Describes each section in the file, helping in the linking and relocation process.
* Symbol and String Tables: Used for name resolution and linking.
* ￼


## DEBUG SYMBOL:

* A debug symbol is a special kind of symbol that attaches additional information to the symbol table of an object file, such as a shared library or an executable. 
* This information allows a symbolic debugger to gain access to information from the source code of the binary, such as the names of identifiers, including variables and routines.
* The symbolic information may be compiled together with the module's binary file, or distributed in a separate file, or simply discarded during the compilation and/or linking. This information can be helpful while trying to investigate and fix a crashing application or any other fault
* Debug symbols typically include not only the name of a function or global variable, but also the name of the source code file in which the symbol occurs, as well as the line number at which it is defined. Other information includes the type of the symbol (integer, float, function, exception, etc.), the scope (block scope or global scope), the size, and, for classes, the name of the class, and the methods and members in it.
* Part of the debug information includes ithe line of code in the source file which defines that symbol (a function or global variable), as well as symbols associated with exception frames.
* This information may be stored in the symbol table of an object file, executable file, or shared library, or may be in a separate file.
* Debugging information can take up quite a bit of space, especially the filenames and line numbers. Thus, binaries with debug symbols can become quite large, often several times the stripped file size.[2] To avoid this extra size, most operating system distributions ship binaries that are stripped, i.e. from which all of the debugging symbols have been removed. This is accomplished, for example, with the strip command in Unix. If the debugging information is in separate files, those files are usually not shipped with the distribution.
* Run from main go build -ldflags "-s -w"  -s ommits the symbol table and  -w which disables DWARF debugging information generation go build -ldflags "-s -w" -o ./bin/cloud/acuetl
