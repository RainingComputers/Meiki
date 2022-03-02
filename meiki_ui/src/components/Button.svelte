<script lang="ts">
    export let onClick: any = undefined
    export let label: string = "Button"
    export let expand: boolean = false
</script>

<button
    on:click={onClick}
    class="bg-blue-500 drop-shadow-card font-bold text-white rounded-xl py-2 px-5"
    class:w-full={expand}
>
    <span class=" text-sm">{label}</span>
</button>
