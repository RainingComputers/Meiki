<script lang="ts">
    import EditableLabel from "$cmp/EditableLabel.svelte"

    export let title: string
    export let onEnter: (text: string) => void
</script>

<span
    class="flex flex-row flex-grow text-gray-200 justify-center font-light items-center"
>
    <EditableLabel text={title} {onEnter} />
    <slot />
</span>
