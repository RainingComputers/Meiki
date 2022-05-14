<script lang="ts">
    import { afterUpdate } from "svelte"
    import Editor from "$cmp/Editor.svelte"
    import Renderer from "$cmp/Renderer.svelte"
    import Logo from "./Logo.svelte"

    export let fontSize = 18

    let showRenderer = false
    let showEditor = false
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

<div class="flex flex-row flex-grow justify-center">
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
    {#if !showRenderer && !showEditor}
        <div class="opacity-5 flex justify-center items-center">
            <Logo width="800em" />
        </div>
    {/if}
</div>
