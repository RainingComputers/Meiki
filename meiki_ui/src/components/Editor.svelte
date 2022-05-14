<script lang="ts">
    import { onDestroy, onMount } from "svelte"
    // import currentNoteText from "$lib/stores/currentNoteText" // TODO: sync with current note text

    export let fontSize: number
    export let editorId: string = "testing" // TODO: change this

    let editor: any = undefined

    onMount(() => {
        // @ts-ignore
        window.ace.config.set("basePath", "/")
        // @ts-ignore
        editor = window.ace.edit(editorId)
        editor.setTheme("ace/theme/xcode")
        editor.session.setMode("ace/mode/markdown")
        editor.setFontSize(fontSize)
        editor.setShowPrintMargin(false)
        editor.setHighlightActiveLine(false)
    })

    onDestroy(() => {
        if (editor) editor.destroy()
    })

    export function focus() {
        // set timeout so the editor does not capture key events like the enter key
        // there is no other way to do this
        setTimeout(() => editor.focus(), 1)
    }

    export function getValue(): string {
        return editor.getValue()
    }

    export function setValue(content: string) {
        editor.setValue(content, 1)
    }
</script>

<div class="w-full h-full" id={editorId} />
