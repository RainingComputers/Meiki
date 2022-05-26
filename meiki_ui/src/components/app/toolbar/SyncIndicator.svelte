<script lang="ts">
    import Badge from "$cmp/toolbar/Badge.svelte"
    import CheckIcon from "$cmp/icons/CheckIcon.svelte"
    import RefreshIcon from "$cmp/icons/RefreshIcon.svelte"

    export let changesNotSaved: boolean = false
    export let show: boolean = false
</script>

{#if show}
    {#if changesNotSaved}
        <Badge label="saving">
            <RefreshIcon />
        </Badge>
    {:else}
        <Badge label="saved">
            <CheckIcon />
        </Badge>
    {/if}
{/if}
