export const load = async ({ params }) => {
    console.log(params);
    console.log(params.school);

    const school = await fetch(`http://localhost:8080/school/${params.school}`);
    return {
        slug: params.school
    }
}