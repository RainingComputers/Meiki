<script lang="ts">
    import Editor from "$cmp/Editor.svelte"
    import Renderer from "$cmp/Renderer.svelte"
    import { showRenderer, showEditor } from "./stores"
</script>

<div class="flex flex-row flex-grow justify-center">
    {#if $showEditor && !$showRenderer}
        <div class=" bg-green-100 w-3/5">
            <Editor />
        </div>
    {/if}
    {#if $showRenderer && !$showEditor}
        <div class=" bg-red-100 w-4/5">
            <Renderer />
        </div>
    {/if}
    {#if $showRenderer && $showEditor}
        <div class=" bg-green-100 flex-1"><Editor /></div>
        <div class=" bg-red-100 flex-1"><Renderer /></div>
    {/if}
</div>
