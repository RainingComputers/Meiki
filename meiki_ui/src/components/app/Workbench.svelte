<script lang="ts">
    import { afterUpdate } from "svelte"
    import Editor from "./Editor.svelte"
    import Renderer from "./Renderer.svelte"
    import Logo from "$cmp/app/Logo.svelte"
    import { currentNote } from "$lib/stores/currentNote"

    export let fontSize = 18
    export let showRenderer = false
    export let showEditor = false

    let editor: Editor
    let splitEditor: Editor

    export function toggleRenderer() {
        showRenderer = !showRenderer
    }

    export function toggleEditor() {
        showEditor = !showEditor
    }

    function focus() {
        if (editor) editor.focus()
        if (splitEditor) splitEditor.focus()
    }

    afterUpdate(focus)
</script>

<div class="flex flex-row flex-grow justify-center items-center">
    <div class="flex flex-row flex-grow justify-center h-full w-full">
        {#if $currentNote}
            {#if showEditor && !showRenderer}
                <div class=" bg-green-100 w-full">
                    <Editor {fontSize} bind:this={editor} />
                </div>
            {/if}
            {#if showRenderer && !showEditor}
                <div class=" bg-red-100 w-4/5">
                    <Renderer />
                </div>
            {/if}
            {#if showRenderer && showEditor}
                <div class=" bg-green-100 flex-1">
                    <Editor {fontSize} bind:this={splitEditor} />
                </div>
                <div class=" bg-red-100 flex-1"><Renderer /></div>
            {/if}
        {/if}
    </div>
    <div class=" opacity-4 fixed">
        <Logo width="700em" />
    </div>
</div>
