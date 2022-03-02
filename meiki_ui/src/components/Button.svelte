<script lang="ts">
    export let onClick: any = undefined
    export let label: string = "Button"
    export let expand: boolean = false
</script>

<button
    on:click={onClick}
    class="bg-blue-600 drop-shadow-card font-bold text-white rounded-xl py-2 px-5 hover:bg-blue-500"
    class:w-full={expand}
>
    <span class=" text-sm">{label}</span>
</button>
