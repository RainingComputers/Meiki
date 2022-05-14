<script>
    import currentNoteText from "$lib/stores/currentNoteText"
    import Renderer from "$cmp/Renderer.svelte"
</script>

<Renderer text={$currentNoteText} />
