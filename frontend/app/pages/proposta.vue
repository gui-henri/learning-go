  <script setup lang="ts">
  import {
    Carousel,
    CarouselContent,
    CarouselItem,
    CarouselNext,
    CarouselPrevious,
  } from '@/components/ui/carousel'
  </script>

  <template>
    <div class="flex items-start justify-center h-screen">
      <Carousel class="w-150 h-150">
        <CarouselContent>
          
          <!-- Slide 1 -->
          <CarouselItem>    
            <div class="flex flex-col items-center">
              <img 
                src="https://cdn.bitrix24.com.br/b22035513/landing/750/750d952c46e82fa4ef81b1103248c565/Atendimento_Domiciliar_-_HomeCare_1x.png" 
                class="w-full h-auto rounded-lg"
              > 
              <p class="mt-2 text-center text-red-700 text-sm">
                <b>Atendimento 24H para nossos pacientes</b>
              </p>

              <p class="mt-2 text-center text-white-700 text-sm text-justify">
                O Home Care engloba serviços de saúde oferecidos em domicílio, que exigem um planejamento especial junto ao paciente e às equipes 
                médicas assistentes, em função das necessidades específicas de cada caso. 
                Os critérios da assistência seguem as normas orientadas pela Associação Brasileira de Home Care, além da avaliação específica do paciente. 
              </p>
            </div>
          </CarouselItem>

          <!-- Slide 2 -->
          <CarouselItem>
            <div class="flex flex-col items-center">
              <img 
                src="https://telemedicinamorsch.com.br/wp-content/uploads/2017/02/iStock-592647828.jpg" 
                class="w-full h-auto rounded-lg"
              >
              <p class="mt-2 text-center text-blue-700 text-sm">
                <b>Tecnologia em medicina</b>
              </p>
              <p class="mt-2 text-center text-white-700 text-sm text-justify">
              O Home Care pode ser dividido em dois formatos: Internação Domiciliar e Assistência Domiciliar.<br>
              Em ambos, o serviço traz os seguintes benefícios ao paciente:

              Conforto
              Humanização e presença constante dos familiares
              Atendimento com equipe multidisciplinar
              Diminuição no risco de infecção hospitalar 
            </p>

            </div>
          </CarouselItem>

          <!-- Slide 3 -->
          <CarouselItem>
            <div class="flex flex-col items-center">
              <img 
                src="https://cdn.prod.website-files.com/647e410aa35e6ce70fed96cd/67449063bc02321b87ce595a_bom-atendimento-medico.jpeg" 
                class="w-full h-auto rounded-lg"
              >
              <p class="mt-2 text-center text-green-700 text-sm">
                <b>Atendimento humanizado com qualidade</b>
                </p>
                <p class="mt-2 text-center text-white-700 text-sm text-justify ">
                Serviço de atendimento do paciente para aplicações de medicamentos, terapias entre outras necessidades específicas por período pré-determinado com objetivo de evolução do paciente.
  Esse serviço atende às necessidades pontuais do paciente. Ele recebe em casa os cuidados da nossa equipe de profissionais, que realiza práticas terapêuticas como a administração de antibióticos, curativos, terapia de oxigênio, entre outras. A Assistência Domiciliar é recomendada nos casos que não exigem internação ou atenção permanente, mas precisam de tratamento especial em momentos distintos.
              </p>
            </div>
          </CarouselItem>

          <!-- Slide 4 -->
      
          <CarouselItem>    
            <div class="flex flex-col items-center">
              
              <img
                src="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wCEAAkGBwgHBgkIBwgKCgkLDRYPDQwMDRsUFRAWIB0iIiAdHx8kKDQsJCYxJx8fLT0tMTU3Ojo6Iys/RD84QzQ5OjcBCgoKDQwNGg8PGjclHyU3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3Nzc3N//AABEIAJQAlAMBEQACEQEDEQH/xAAcAAEAAgMBAQEAAAAAAAAAAAAABgcBBAUCAwj/xABBEAABAwICBgYHBgMJAQAAAAABAAIDBAUGEQcSITFBURMiYXGBwRQjQpGhsdEyNlJidOFTgpIWMzRDRGNyovEV/8QAGwEBAAIDAQEAAAAAAAAAAAAAAAQFAQMGAgf/xAAyEQACAgIABQIFAgUFAQAAAAAAAQIDBBEFBhIhMTJBEyIzUXFhgRRCobHBJJHR8PEj/9oADAMBAAIRAxEAPwC8UAQBAYQwYzCGSOX7GdlsutHPU9LP/BhGs7x4DxWmy+EPJY4nCcrKfyR7fqQS7aVLjMS210kNOzg+Xrv8gPioc81/ynSY3K8Et3S3+CM1uLL9XawnutQGu3tjdqD4KPLIsfuXNPBsKrxBP8nKkqaiUnpaiaTP8cjj8ytTk35ZOhRXX6Y6PmHuac2udn2ErybHFPyjcprvcqRwdS3CqjI4CZ2XuzWyNs4+GRbMDGt9UEzu0GkPEVEW61VHVN4tqI8/iMit0cuxFZfy5h2ela/BMbNpSoJy1l2p30jzvkZm9g8wpVeZGXq7FBlctZFe3U+pE7oa6lr4Gz0dRHNG7aHMOalqSfdHP2U2VPpmtM2QV6NZlAEAQBAEAQGM0Bzr1eqGy0jqm4TNjYPsjPNzzyA4leJzjBbZIxsW3Jn0VLbKhxRj+5Xcvgoi6iozsyY7rvHaeHcFW3ZUpdkdvw7l+mjU7vml/Qh+fHiobezoVFJaRhD0EAQBAEAWAZByWQblqutdaKkVFuqJIZOOqdju8cVshbKD2mRMrBpyY9Nkd/3LWwhpDpbm5lJdgylqzsa/PqSd3I9isqcpT7PycRxPgNuK+qrvEnjXa3/qlruc+ekAQBAEBgoDg4sxNS4cojNOQ+d4IhgB2vPkO1abbVWtsnYGBbm29EF292UffLzW3utdV18pc4/YYD1WDkAqiy2Vj2z6NhYNWJX0QX7/AHOatZOQWAEAQBAEAQBAEAWQFlGHHa7liYCx4+idHbL1KX0x2RVL3ZmPkHHl28FPxsn+WRyHGeBp7ux1+UW2xwc0OBBB3EKxON1rsekAQGEBzr9dqay2yauq3ZMjGwcXHgB2rxZNQW2SMXGnlWquHllA328VV7uMtZWOzc45MYN0beACpbbHY9s+mYOFViVKEP3/AFOctROCAIAgCAIAvRhyS7sm+HdHNyusDKmtlFFA/aGubrSEd3BS68RyW5HNZ3MVVEnCpdT/AKEik0T0XRerulSH83NaR7lveDD7lVHmjIUu8VohOJ8HXPDo6ScNnpSchPEDkO8cFFtxpV9/Y6Dh3GaMz5fTL7EcUYu0EAQwyz9GOLS5zLJcpMyf8NK8/wDQ+SssW/fys4rj/CVX/qKl291/ktEKeckjKGTydyApHSViI3e8OpKd5NHSOLAAdj38T5KpyruqWl4O/wCX+HfAp+LNfNL+xDlEOj0FgBAEAQBAFkE00XWGO7Xx9VVM1qejaHBp3GQnZ7tvwUzEqUpNs5rmLOlRSq4PTl/YutrcgrU4IzkgPhV00VVTyU87A+KQFrmu4hYa2mmeoTlXJSj2aPz1ie1Gy32roNupG/1RPFh2j4Kkur6JaR9Q4Xl/xWLGx+TlrUWAQHqOR8T2yRuLHtObXA7QVlPT2jxOCnFxl4ZfeBr+MQWSOd+XpMXq5x+bLf4q6osU4nzHiuC8PIcNdn3X4JHmt5Wtkbx5ev8A4uHKiaN3r5R0UP8AyP0GZWi+zohssuFYjysqMPbyUIfHxVIz6fFaRhYPQQBAEAQBAEBauheVnolzi9vpWuy7MsvIq0wX8rRw/NMGroS9tFlhTjlTKA8lAUhpVkZJi+UMyzZCxru/aqnN+ofQeWotYXf7sh6iHQ+wWAFkEt0a3o2rEcUMjsqes9U/sd7J9+zxUvEs6Za+5z/MOH8fG615iXjmrY+ebKi0w3IzXemtrD1KeLpHgfid+w+Krc2ffpO25XxkqpXP37FfKAdaFgBAEAQBAEAWQd7Bd/OHb02qIc6nkHRztH4eY7Qt9Fvw5bKji/D/AONo6V6l4L5oa6mrqZlRSStlieM2vadhVxGSkto+b2VTqk4zWmbGY3r0eNnHxPiCisFvfUVUg1yPVRA9aQ8gtdlkYR7kzCwrMyxQgigbjWS3CuqKyoOcs7y93iqScnKW2fTsWiOPTGuPsay8EgIAgPUb3RyNkjOT2nWaeRG5ek9PZ4sgrIuL9z9F4frxdLJQ1wyzmha5wz3HLaPeryt9cUz5Pl0/BvnX9mUZjOr9NxRcps8x0xY3ubs8lU5Et2M+jcHp+HhQX3WzirQWoWAEAQBAEAQBABsKyYa2dC03u5WeTXttZLBnvYHZtP8AKdi2QtnDwyHk8Ox8n6kUztyaQ8SyR6npcbPzMiAK3PLtaK2PLuEnvTZHKytqq+oNRWzyVEx9uRxcf2UeU5TfcuKMaqiPTVFJGuTmvJvCwZCGAgCyC39GN0b/AGXbDI4ZwzvZtPDYfNWuNbqtHz7j2NrNk17pMqarkMtVPIduvI52feSVWSe22d5RD4dagvY+K8m0LACAIAgCAIBxWTDaS2di04YvN3cPQrfMYz/myNLWe8rdCicvYrcni+Jj+qe/wTezaKjm2S81mY/g04y97ipleEl3kzncrmeUu1MdfqyUzYCw5JRimFtYzIf3rCQ/+repH8PXrWiljxnNjPr63/gh940V1Mes+z1jZW8Ipxk7+oKLPC94svsXmhdlfH90Qm6WG62l5bcKCeID2yzNh/mGxRJUzj5R0WNxLGyPRNHNWon7CwAgAWQzs2W6voKaSNj3AOkLthy4AeS2ws6VorcrCjfZ1tHIkBa9zTvBIK1FhF7SZ5Q9BYAQBAEAQHQsUNvnuLGXeofT0mqS97BmdnAd621KLfzeCFnzyIVbx47kTS34kwPZph6DZquZwyzqJGBzj29Z3kFLjbRB9kc5dw7i2THdliX6f+IsyyXOjvFBHWW9+vA/MDZkWkcCOBU+EoyW0cpk49uPa67fKOiF7NAQBAcXE19t1hoemuRLmvzayJrQ50h5AFarbIwXzEzCw7sqzoq8lb1l4wLepj6Va6ygld/qImho8Q0n5KE7MeflaOorwuMYkfkmpL7f9RCrlHTRV08dFMZqZryI5CNrhzUKelLSOmxZ2TqjK1al7mqvJICA2KeF8jSWjMa2Syls1ynGL0z7XunNJeK6ncMjHUPHhrHJe7VqbRHwLPi40J/dGitZMCwAgCAIAgGayDOabGvYtbQvM51HcoSeq2RrgOWYy8lZ4T3FnDc0QSuhJe6LKU45YIAgKf0yTOfeqGEnqx05IHaT+yrc5vaR23K1a+HORXygHV6QO1DOgsABAT3AdhNwtE1Q4HI1Dg3uDW+eanUVdUNs5PjGc6cnoX2RzdJ1F6Fi6pdq5NqWNmaeeew/ELxlx1YTeXb/AImGov27EUUUvwsAIAgCAIAgCGS0NCu6698fmrPB8M4nmr11/gtFTzkggMFAU3ph+8dP+mHzKrM71I7rlb6E/wAkDUE6gLACAd+5ZQ2Xzo8oDRYRoGuGq6VnTOB/Nt+WSuseGq0j5jxa/wCLmTl+uv8AY4el+1ekWqC5RtzdSv1ZCPwO2fPJacyvqj1fYsOWsv4WQ6m+0iolVnfBYAQBAEAQBAEBaGhXdde+PzVng+GcVzX66/3LRCnnJBAYO5AU3ph+8lP+lHzKrM71I7rlb6E/yQNQTqAsALIOhYba+7XqjoGAnppAHZcG73H3ZrbTDrmkQeI5Kx8eVn6H6LijbFEyNjQGMaA0cgrxLSPlcptvf3PncKKGvoZ6SdodHMwscD2ry4ppp+57qsdU1OPlH51vVsns9zqKCpB14nkAkfabwPiFSW19EtH1PByo5NEbIvyaS1kwLACAIAgCAIC0NCu6698fmrPB8M4rmv11/uWiFPOSCAwdyApvTD95Kf8ASj5lVmd6kd1yt9Cf5IGoJ1AWAFkMtTRDYnRxTXqpZkZPV04I9ni7x3eCs8OrS6mcNzJnKc1jwfjuyzNUFTjlgRsQEE0nYXN0ohc6NmdZTN6zRvkZy7womVT1raOg4DxL+Ft+FP0y/oymyMlUvsfQYtNGEMhYAQBAEAQFoaFd1174/NWeD4ZxXNfrr/ctEKeckEBg7kBTemH7yU/6UfMqszvUjuuVvoT/ACQNQTqAgO3hKwTYhuzKRmbYG9aeQey36lbqKviSKrivEI4dTl/M/Bf1HTRUdPFTwMDIomhrGjgArpLS0fNJzlZJzl5ZsLJ5CA8uGaAqXSNgt1JJLd7VFnTuJdURN9g8XAcuarsnH/midlwPjPZY9z/DK75KvOvTMIZCwAgCAIC0NCu6698fmrPB8M4vmv11/uWip5yIQGDuQFN6YfvJT/pR8yqzO9SO65W+hP8AJA1BOoN6z2qrvNeyioI9eZ/Hg0cyeS2V1ux6RDzM2vEr+JYXxhbD1Lh22NpabJz3daWUjbI7mrmqtVx0j5rnZtmXa5zO0thDMoAgCA8uY1wOsMweBQLt4Kwxro8MjpK+ws6x60lLnkD2s+igZGLv5onV8J4+60qcjx9/+SspY3RSOjka5j2nJzXDIgquaaemdpCyNi6ovaPCwewsAIAgLR0K/ZuvfH5qzwfDOL5r+pX+5aCnnIhAYO5AU3ph+8lP+mHzKrM31I7nlf6E/wAkew1hi4YinDaOPVgB9ZO/7LfqewKPVRKx9i04hxanDj3e5fYuvDOHKHD1F6PRtLnu2yzO+089v0VvXVGtaR8+zc63Ms67GdoDYthDS0ZQBAEAQBAeSEMaI/iPCFrv7C6pi6Kpy2VEWQcO/n4rTZRGaLHC4nkYcvkfb7PwVfftH16tZL6aP06nHtQ7HDvb9FX2Yko+O52GHzFjXdrPlZEpI3xPMcrHMeN7XDIjwUVxa8l9XbCa3F7PO7esGwIC2dDFO5tuuE5aQHzBrTzyH7q0wlqLZwnNFieRCK9kWQppzAQGCgI3fcIUF9vEFfcC97YY9QQDY123Pb9FpnRGyW2WOLxO7FqlXV7+53qWlhpIWQ00TYomDJrGDIBbUkuyIE5ynLqk9s+yyeTKAIAgCAIAgCAIDw7fksnnejQuNltl0GrcKGnqORkjBI7jvC8ShGa7okVZd9L/APnJoid30eWDo3yQx1EDv9uXYPA5qLZjVl1i8ezV2ct/lEBuGHqWlrWwsmqHNLsusW57+5Q5VRTOmo4hdZV1S0XPhu2UtotMFJRMLYmt1tpzLidpJKs6oqMNI4LMyLMi+U7PJ1QtpGCAIAgCAIAgCAID/9k="
                class="w-50 h-50 rounded-lg"
                >

<p class="mt-2 text-sm text-red-700 text-justify"><b>INTERNE SOLUÇÕES EM SAÚDE, A ESCOLHA CERTA PARA HOME CARE</b></p>

            <div class="mt-2 text-sm text-white-700 whitespace-pre-line text-justify">
              <ul class="space-y-3 justify-middle">
    <li>O Home Care engloba serviços de saúde oferecidos em domicílio.</li>
  <li>Técnicos qualificados.</li>
  <li>Atendimento humanizado.</li>
  <li>Colaboradores especializado.</li>
  <li>Sistemas internos ágeis e seguros.</li>
  <li>Espaço amplo para atendimento.</li>
  <li>Equipamentos eficientes.</li>
  <li>Focada no bem-estar de seus pacientes.</li>



    <li><b>Saiba mais em : </b><a href="interne.com.br" class="text-blue-500 underline">
      Nosso site parceiro.
    </a>
 </li>
    </ul>
    




  </div>
            </div>
          </CarouselItem>
          <CarouselItem>

            
          </CarouselItem>


        </CarouselContent> 
        <CarouselPrevious />
        <CarouselNext />
      </Carousel>
    </div>
  </template>
