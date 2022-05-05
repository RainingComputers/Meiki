<script lang="ts">
    import currentNoteText from "$lib/stores/currentNoteText"
</script>

<textarea
    class="w-full h-full resize-none p-2"
    data-cy="editor"
    bind:value={$currentNoteText}
/>
