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
    class=" hover:bg-blue-500 hover:text-white text-slate-900 stroke-black hover:stroke-white"
    on:click={onClick}
>
    <div
        class="flex flex-row px-4 gap-3 py-1 items-center"
        class:bg-blue-500={isButtonChecked}
        class:text-white={isButtonChecked}
        class:stroke-white={isButtonChecked}
    >
        <span class="w-5 h-5 flex-shrink-0">
            <ItemsIcon />
        </span>
        <span>{title}</span>
    </div>
</div>
