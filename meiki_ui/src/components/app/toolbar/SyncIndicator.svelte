<script lang="ts">
    import Badge, { BadgeType } from "$cmp/toolbar/Badge.svelte"
    import CheckIcon from "$cmp/icons/CheckIcon.svelte"
    import RefreshIcon from "$cmp/icons/RefreshIcon.svelte"
    import WarningIcon from "$cmp/icons/WarningIcon.svelte"

    export let changesNotSaved: boolean = false
    export let show: boolean = false
    export let error: string = ""
</script>

{#if show}
    {#if error}
        <Badge label={error.toUpperCase()} type={BadgeType.ERROR}>
            <WarningIcon />
        </Badge>
    {:else if changesNotSaved}
        <Badge label="">
            <RefreshIcon />
        </Badge>
    {:else}
        <Badge label="">
            <CheckIcon />
        </Badge>
    {/if}
{/if}
