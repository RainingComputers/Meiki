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
    import { afterUpdate, createEventDispatcher } from "svelte"
    import Editor from "$cmp/Editor.svelte"
    import Renderer from "$cmp/Renderer.svelte"
    import Logo from "$cmp/app/Logo.svelte"

    export let fontSize = 18
    export let showEditorAndRenderer: boolean

    let text = ""
    let editor: Editor
    let showRenderer = false
    let showEditor = false
    const dispatchEvent = createEventDispatcher()

    function onEditorChange() {
        text = editor.getValue()
        dispatchEvent("textChange", { text })
    }

    function focus() {
        if (editor) editor.focus()
    }

    export function toggleRenderer() {
        showRenderer = !showRenderer
    }

    export function toggleEditor() {
        showEditor = !showEditor
    }

    export function setText(newText: string) {
        text = newText
        editor.setValue(newText)
    }

    afterUpdate(focus)
</script>

<div class="flex flex-grow justify-center items-center">
    <div class="flex justify-center h-full w-full">
        {#if showEditorAndRenderer}
            <div class={getEditorClass(showEditor, showRenderer)}>
                <Editor
                    bind:this={editor}
                    {fontSize}
                    onChange={onEditorChange}
                    editorId="workbenchEditor"
                    initialText={text}
                />
            </div>
            <div class={getRendererClass(showEditor, showRenderer)}>
                <Renderer {text} />
            </div>
        {/if}
    </div>
    <div class=" opacity-4 fixed">
        <Logo width="900em" />
    </div>
</div>
