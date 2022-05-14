<script lang="ts">
    import { onDestroy, onMount } from "svelte"
    import currentNoteText from "$lib/stores/currentNoteText"
    import currentNote from "$lib/stores/currentNote"

    export let fontSize: number

    const editorId: string = $currentNote
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
        editor.setValue($currentNoteText, 1)
        editor.getSession().on("change", () => {
            currentNoteText.set(editor.getValue())
        })
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
