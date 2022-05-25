<script lang="ts">
    import ItemsIcon from "$cmp/icons/ItemsIcon.svelte"
    import Watermark from "$cmp/explorer/watermark/Watermark.svelte"

    export let message: string
</script>

<Watermark {message}>
    <ItemsIcon />
</Watermark>
