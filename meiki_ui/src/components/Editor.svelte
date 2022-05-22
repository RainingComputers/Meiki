<script lang="ts">
    import { onDestroy, onMount } from "svelte"

    export let fontSize: number
    export let editorId: string
    export let onChange: () => void
    export let initialText: string

    let aceEditor: any = undefined

    onMount(() => {
        // @ts-ignore
        window.ace.config.set("basePath", "/")
        // @ts-ignore
        aceEditor = window.ace.edit(editorId)
        aceEditor.setTheme("ace/theme/xcode")
        aceEditor.session.setMode("ace/mode/markdown")
        aceEditor.setFontSize(fontSize)
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
