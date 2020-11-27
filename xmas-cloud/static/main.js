require.config({ paths: { vs: '/monaco-editor/min/vs' } });

require(['vs/editor/editor.main'], function () {

    // validation settings
    monaco.languages.typescript.javascriptDefaults.setDiagnosticsOptions({
        noSemanticValidation: false,
        noSyntaxValidation: false
    });

    // compiler options
    monaco.languages.typescript.javascriptDefaults.setCompilerOptions({
        target: monaco.languages.typescript.ScriptTarget.ES5,
        moduleResolution: monaco.languages.typescript.ModuleResolutionKind.NodeJs,
        module: monaco.languages.typescript.ModuleKind.CommonJS,
        lib: ["es5"],
        allowNonTsExtensions: true
    });

    // extra libraries
    var libSource = [
        'interface Console {',
        '    log(...data: any[]): void;',
        '    error(...data: any[]): void;',
        '    warn(...data: any[]): void;',
        '}',
        'declare var console: Console;',
        'declare function require(path: string): any;',
        '/**',
        '* Retrieve XMAS.',
        '*/',
        'declare function xmas(): string;'
    ].join('\n');
    var libUri = 'ts:runtime.d.ts';
    monaco.languages.typescript.javascriptDefaults.addExtraLib(libSource, libUri);

    var libSource = [
        'interface OSFunctions {',
        '    /**',
        '     * Execute a system command.',
        '     */',
        '    static exec(cmd: string, ...args: string[]):string',
        '    /**',
        '     * Lists files.',
        '     */',
        '    static ls(path: string):string[]',
        '    /**',
        '     * Read file.',
        '     */',
        '    static readFile(path: string, base64?: boolean):string',
        '    /**',
        '     * Write file.',
        '     */',
        '    static writeFile(path: string, content:string):boolean',
        '}',
        'declare var OS: OSFunctions;'
    ].join('\n');
    var libUri = 'ts:OS.d.ts';
    monaco.languages.typescript.javascriptDefaults.addExtraLib(libSource, libUri);
    // When resolving definitions and references, the editor will try to use created models.
    // Creating a model for the library allows "peek definition/references" commands to work with the library.
    monaco.editor.createModel(libSource, 'typescript', monaco.Uri.parse(libUri));

    const libDemoCode = `/* XMAS CLOUD Example Code */

function logJSON(data) {
    console.log("\\n\`\`\`json\\n" + JSON.stringify(data, null, 3) + "\\n\`\`\`\\n");
}

console.log("# Hello Advent\\n");

console.log(xmas() + "\\n");
console.log(this + "\\n");

logJSON({a: 1, b: "test"});`;

    const savedCode = window.localStorage.getItem("code");

    var editor = monaco.editor.create(document.getElementById('editor'), {
        value: savedCode ?? libDemoCode,
        language: "javascript",
    
        roundedSelection: false,
        scrollBeyondLastLine: false,
        readOnly: false,
        theme: "vs-dark",
        lineNumbers: "on",
        automaticLayout: true
    });

    const defaultOutput = `# Help

You can run the code in the editor with the key F9.

## Shortcuts

F1: Show command palette
F3: Search
Ctrl+Space: Autocompletion
F9: Run code`;
    

    var output = monaco.editor.create(document.getElementById('output'), {
        value: defaultOutput,
        language: "markdown",
    
        readOnly: true,
        theme: "vs-dark",
        lineNumbers: "off",
        automaticLayout: true,
        wordWrap: "on"
    });

    editor.getModel().onDidChangeContent((event) => {
        window.localStorage.setItem("code", editor.getValue());
    });

    var runCode = function() {
        output.setValue("Running script...")

        fetch('/api/exec', {
            method: "POST",
            headers: {
                'Content-Type': 'application/javascript'
            },
            body: editor.getValue()
        })
            .catch(e =>  output.setValue(`Execution failed:\n ${e}`))
            .then(response => response.text())
            .then(text => output.setValue(text));
    };

    editor.addAction({
        // An unique identifier of the contributed action.
        id: 'my-run-code',
    
        // A label of the action that will be presented to the user.
        label: 'Run Code',
    
        // An optional array of keybindings for the action.
        keybindings: [
            monaco.KeyCode.F9
        ],
  
        contextMenuGroupId: 'navigation',
    
        contextMenuOrder: 1.0,

        run: runCode
    });

    output.addCommand(monaco.KeyCode.F9, runCode);
    editor.focus();
});

