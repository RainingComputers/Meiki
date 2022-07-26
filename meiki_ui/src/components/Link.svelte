<script lang="ts">
    export let text: string
    export let target: string = "#"
</script>

<a class=" text-background-primary font-bold" href={target}>{text}</a>
