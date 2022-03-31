<script lang="ts">
    export let showOverlay: Boolean = false

    export function onClick() {
        showOverlay = false
    }

    export function showModal() {
        showOverlay = true
    }
</script>

<div
    class="fixed inset-0 bg-gray-600 bg-opacity-60 overflow-y-auto h-full w-full flex flex-col align-center justify-center"
    class:hidden={!showOverlay}
    class:z-10={showOverlay}
    data-cy="modalOverlay"
    on:click={onClick}
>
    <div on:click|stopPropagation>
        <slot />
    </div>
</div>
