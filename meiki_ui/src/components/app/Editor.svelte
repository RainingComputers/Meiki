<script lang="ts">
    import Editor from "$cmp/Editor.svelte"
    import currentNoteText from "$lib/stores/currentNoteText"
    import currentNote from "$lib/stores/currentNote"

    export let fontSize: number

    let editor: Editor

    function onChange() {
        currentNoteText.set(editor.getValue())
    }

    export function focus() {
        editor.focus()
    }
</script>

<Editor
    {fontSize}
    {onChange}
    editorId={$currentNote}
    initialText={$currentNoteText}
    bind:this={editor}
/>
