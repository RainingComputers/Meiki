<script lang="ts">
    import { onDestroy, onMount } from "svelte"

    import * as ace from "brace"
    import "brace/mode/markdown"
    import "brace/theme/textmate"

    export let fontSize: number
    export let editorId: string
    export let onChange: () => void
    export let initialText: string

    let aceEditor: ace.Editor

    onMount(() => {
        aceEditor = ace.edit(editorId)
        aceEditor.setTheme("ace/theme/textmate")
        aceEditor.session.setMode("ace/mode/markdown")
        aceEditor.setFontSize(fontSize + "px")
        aceEditor.setShowPrintMargin(false)
        aceEditor.setHighlightActiveLine(false)
        aceEditor.setValue(initialText, 1)
        aceEditor.getSession().on("change", onChange)
    })

    onDestroy(() => {
        if (aceEditor) aceEditor.destroy()
    })

    export function focus() {
        // set timeout so the editor does not capture key events like the enter key
        // there is no other way to do this
        setTimeout(() => aceEditor.focus(), 1)
    }

    export function getValue(): string {
        return aceEditor.getValue()
    }

    export function setValue(content: string) {
        aceEditor.setValue(content, 1)
    }
</script>

<div data-cy="editor" class="w-full h-full" id={editorId} />
