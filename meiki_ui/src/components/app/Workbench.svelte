<script context="module" lang="ts">
    function getEditorClass(showEditor: boolean, showRenderer: boolean) {
        if (showEditor && showRenderer) return "bg-green-100 flex-1"
        if (showEditor && !showRenderer) return "bg-green-100 w-full"

        return "hidden"
    }

    function getRendererClass(showEditor: boolean, showRenderer: boolean) {
        if (showEditor && showRenderer) return "bg-red-100 flex-1"
        if (showRenderer && !showEditor) return "bg-red-100 w-3/4"

        return "hidden"
    }
</script>

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

    export function toggleRenderer() {
        showRenderer = !showRenderer
    }

    export function toggleEditor() {
        showEditor = !showEditor
    }

    function focus() {
        if (editor) editor.focus()
    }

    afterUpdate(focus)
</script>

<div class="flex flex-grow justify-center items-center">
    <div class="flex justify-center h-full w-full">
        {#if $currentNote}
            <div class={getEditorClass(showEditor, showRenderer)}>
                <Editor {fontSize} bind:this={editor} />
            </div>
            <div class={getRendererClass(showEditor, showRenderer)}>
                <Renderer />
            </div>
        {/if}
    </div>
    <div class=" opacity-4 fixed">
        <Logo width="900em" />
    </div>
</div>
