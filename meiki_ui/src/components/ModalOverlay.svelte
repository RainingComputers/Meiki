<script lang="ts">
    import { scale } from "svelte/transition"

    export let showOverlay: Boolean = false

    export function onClick() {
        showOverlay = false
    }

    export function showModal() {
        showOverlay = true
    }

    export function closeModal() {
        showOverlay = false
    }
</script>

{#if showOverlay}
    <div
        class="fixed inset-0 bg-gray-600 bg-opacity-60 h-full w-full flex flex-col align-center"
        class:hidden={!showOverlay}
        class:z-10={showOverlay}
        data-cy="modalOverlay"
        on:click={onClick}
    >
        <span
            on:click|stopPropagation
            class="max-w-fit m-auto"
            transition:scale={{ duration: 70, opacity: 1 }}
        >
            <slot />
        </span>
    </div>
{/if}
