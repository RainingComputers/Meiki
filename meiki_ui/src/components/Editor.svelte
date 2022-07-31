<script lang="ts">
    import { theme } from "$lib/stores/theme"
    import type AceAjax from "brace"
    import { onDestroy, onMount } from "svelte"

    export let fontSize: number
    export let editorId: string
    export let onChange: () => void
    export let initialText: string

    let aceEditor: AceAjax.Editor

    onMount(async () => {
        const ace = (await import("brace")).default
        ;(await import("brace/mode/markdown")).default
        ;(await import("brace/theme/textmate")).default
        ;(await import("brace/theme/tomorrow_night_bright")).default

        aceEditor = ace.edit(editorId)
        aceEditor.session.setMode("ace/mode/markdown")
        aceEditor.setFontSize(fontSize + "px")
        aceEditor.setShowPrintMargin(false)
        aceEditor.setHighlightActiveLine(false)
        aceEditor.setValue(initialText, 1)
        aceEditor.getSession().on("change", onChange)
        aceEditor.getSession().setUseWrapMode(true)

        const setTheme = (theme: string) => {
            const editorTheme = theme === "light" ? "textmate" : "tomorrow_night_bright"
            aceEditor.setTheme(`ace/theme/${editorTheme}`)
            aceEditor.container.style.background = "transparent"
        }

        theme.subscribe(setTheme)
    })

    onDestroy(() => {
        if (aceEditor) aceEditor.destroy()
    })

    export function focus() {
        // set timeout so the editor does not capture key events like the enter key
        // there is no other way to do this
        setTimeout(() => {
            if (aceEditor) aceEditor.focus()
        }, 1)
    }

    export function getValue(): string {
        return aceEditor.getValue()
    }

    export function setValue(content: string) {
        aceEditor.setValue(content, 1)
    }
</script>

<div data-cy="editor" class="w-full h-full" id={editorId} />
