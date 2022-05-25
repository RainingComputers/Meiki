<script lang="ts">
    import WarningIcon from "$cmp/icons/WarningIcon.svelte"
    import Watermark, {
        WatermarkType,
    } from "$cmp/explorer/watermark/Watermark.svelte"

    export let message: string
</script>

<Watermark type={WatermarkType.ERROR} {message}>
    <WarningIcon />
</Watermark>
