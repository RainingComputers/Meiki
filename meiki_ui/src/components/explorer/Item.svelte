<script lang="ts">
    import ItemsIcon from "../icons/ItemsIcon.svelte"
    export let title: string

    let isButtonChecked: boolean = false

    function onClick() {
        isButtonChecked = !isButtonChecked
    }

    export function isChecked() {
        return isButtonChecked
    }
</script>

<div
    class=" hover:bg-blue-500 hover:text-white text-slate-900 stroke-black hover:stroke-white "
    on:click={onClick}
>
    <div
        class="flex flex-row px-4 gap-3 py-1 items-center   "
        class:bg-blue-600={isButtonChecked}
        class:text-white={isButtonChecked}
        class:stroke-white={isButtonChecked}
    >
        <span class=" h-4 w-4">
            <ItemsIcon />
        </span>
        <span class=" font-light text-sm">{title}</span>
    </div>
    <!-- TODO: Fix styling  -->
</div>
