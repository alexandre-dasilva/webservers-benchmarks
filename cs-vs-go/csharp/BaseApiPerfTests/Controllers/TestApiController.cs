using System;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;

namespace BaseApiPerfTests.Controllers
{
    [ApiController]
    [Route("")]
    public class TestApiController : ControllerBase
    {
        private readonly ILogger<TestApiController> _logger;

        public TestApiController(ILogger<TestApiController> logger)
        {
            _logger = logger;
        }

        [HttpGet("{numberOfAwaits:int?}")]
        public async Task<string> Get(int numberOfAwaits = 0)
        {
            for (var i = 0; i < numberOfAwaits; i++)
            {
                await Task.Run(async () => await Task.Delay(100));
            }

            return Request.Path;
        }
    }
}