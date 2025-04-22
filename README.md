# Generating Software Bill of Materials (SBOM) for Polyglot Applications

## Table of Contents
- [Overview](#overview)
- [Build Tools](#build-tools)
  - [Java](#java)
    - [Maven](#maven)
    - [Gradle](#gradle)
  - [Node.js](#nodejs)
    - [NPM](#npm)
    - [Yarn](#yarn)
  - [Python](#python)
    - [PIP](#pip)
  - [Go](#go)
- [SCA Tools - CI/CD Integration](#sca-tools-ci-cd-integration)
  - [Syft](#syft)
  - [CDXGen](#cdxgen)
- [References](#references)

## Overview

This guide explains how to generate Software Bill of Materials (SBOM) for polyglot applications using CycloneDX plugins and tools for various build systems. An SBOM provides a detailed inventory of all components, libraries, and dependencies used in your software project, which is essential for security and compliance.

## Build Tools

### Java

#### Maven

Add the CycloneDX Maven plugin to your `pom.xml`:

```xml
<plugin>
    <groupId>org.cyclonedx</groupId>
    <artifactId>cyclonedx-maven-plugin</artifactId>
    <version>2.9.1</version>
    <configuration>
        <projectType>application</projectType>
        <outputFormat>json</outputFormat>
        <outputName>application.cdx</outputName>
        <schemaVersion>1.6</schemaVersion>
    </configuration>
</plugin>
```

Maven Key Configuration Options

| Option | Description | Default |
|--------|-------------|---------|
| `projectType` | Specifies the type of project (`application` or `library`) | `application` |
| `outputFormat` | Format of the SBOM (`json` or `xml`) | `json` |
| `outputName` | Name of the output SBOM file | `application.cdx` |
| `schemaVersion` | CycloneDX schema version | `1.6` |

##### Spring Boot - Normal Project

N/A (No additional config requried)

##### Spring Boot - GraalVM Native Project

The `spring graalvm native` project uses the following configuration in its `pom.xml`:

```xml
<plugin>
    <groupId>org.graalvm.buildtools</groupId>
    <artifactId>native-maven-plugin</artifactId>
    <configuration>
        <buildArgs combine.children="append">
            <buildArg>--enable-sbom=classpath,export</buildArg>
        </buildArgs>
    </configuration>
</plugin>
```

##### Quarkus

The `quarkus` project uses the following configuration in its `pom.xml`:

```xml
<plugin>
    <groupId>io.quarkus</groupId>
    <artifactId>quarkus-maven-plugin</artifactId>
    <version>${quarkus.platform.version}</version>
    <extensions>true</extensions>
    <executions>
        <execution>
            <goals>
                <goal>build</goal>
                <goal>dependency-sbom</goal>
            </goals>
            <phase>package</phase>
            <configuration>
                <attachSboms>true</attachSboms>
                <format>json</format>
                <schemaVersion>1.6</schemaVersion>
            </configuration>
        </execution>
    </executions>
</plugin>
```

##### Maven Generating the SBOM

To generate the SBOM, run:

```bash
mvn cyclonedx:makeAggregateBom
```

The SBOM will be generated in the `target` directory with the specified output name.

#### Advanced Configuration

##### Adding Metadata

Both plugins support adding additional metadata to the SBOM:

```xml
<configuration>
    <organizationalEntity>
        <name>Your Organization</name>
        <url>https://your-org.com</url>
        <contact>
            <name>Contact Name</name>
            <email>contact@your-org.com</email>
        </contact>
    </organizationalEntity>
</configuration>
```

##### Adding License Information

You can specify license information for your project:

```xml
<configuration>
    <licenseChoice>
        <license>
            <name>Apache License 2.0</name>
            <url>https://www.apache.org/licenses/LICENSE-2.0</url>
        </license>
    </licenseChoice>
</configuration>
```

#### Gradle

Add the CycloneDX Gradle plugin to your `build.gradle` or `build.gradle.kts`:

```groovy
plugins {
    id 'org.cyclonedx.bom' version '2.2.0'
}

tasks.named('cyclonedxBom') {
    schemaVersion = "1.6"
    includeConfigs = ["runtimeClasspath", "compileClasspath"]
    skipProjects = [rootProject.name]
    projectType = "application"
    includeBomSerialNumber = true
    includeLicenseText = false
    destination = file("build/reports")
    outputName = "application.cdx"
    outputFormat = "json"
}
```

Gradle Key Configuration Options

| Option | Description | Default |
|--------|-------------|---------|
| `schemaVersion` | CycloneDX schema version | `1.6` |
| `includeConfigs` | List of configurations to include | `["runtimeClasspath"]` |
| `skipProjects` | Projects to exclude | `[]` |
| `projectType` | Type of project | `application` |
| `includeBomSerialNumber` | Include unique identifier | `true` |
| `includeLicenseText` | Include full license text | `false` |
| `destination` | Output directory | `build/reports` |
| `outputName` | Output file name | `application.cdx` |
| `outputFormat` | Output format | `json` |

##### Spring Boot - Normal Project

N/A (No additional config requried)

##### Spring Boot - GraalVM Native Project

The `spring graalvm native` project uses the following configuration in its `build.gradle`:

```groovy
graalvmNative {
    binaries {
        configureEach {
            buildArgs.add("--enable-sbom=classpath,export")
            imageName = "sbom-polyglot-java-cdx-gradle-native"
        }
    }
}
```

##### Quarkus

N/A (No support like quarkus-maven-plugin)

##### Gradle Generating the SBOM

To generate the SBOM, run:

```bash
./gradlew cyclonedxBom
```

The SBOM will be generated in the specified destination directory (default: `build/reports`).

#### Advanced Configuration

##### Adding Metadata

Both plugins support adding additional metadata to the SBOM:

```groovy
cyclonedxBom {
    organizationalEntity { oe ->
        oe.name = 'Your Organization'
        oe.url = ['https://your-org.com']
        oe.addContact(organizationalContact)
    }
}
```

##### Adding License Information

You can specify license information for your project:

```groovy
cyclonedxBom {
    licenseChoice { lc ->
        def license = new License()
        license.setName("Apache License 2.0")
        license.setUrl("https://www.apache.org/licenses/LICENSE-2.0")
        lc.addLicense(license)
    }
}
```

### Node.js

#### NPM

Install the CycloneDX Node.js module globally or as a development dependency:

```bash
npm install -g @cyclonedx/bom
```

Or add it to your project:

```bash
npm install --save-dev @cyclonedx/bom
```

##### Generating the SBOM

To generate the SBOM, run:

```bash
cyclonedx-bom -o sbom.json
```

The SBOM will be generated in the current directory with the name `sbom.json`.

#### Advanced Configuration

##### Adding Metadata

You can add metadata to the SBOM using the CycloneDX Node.js module:

```json
{
  "organizationalEntity": {
    "name": "Your Organization",
    "url": "https://your-org.com",
    "contact": {
      "name": "Contact Name",
      "email": "contact@your-org.com"
    }
  }
}
```

##### Adding License Information

You can specify license information for your project:

```json
{
  "license": {
    "name": "Apache License 2.0",
    "url": "https://www.apache.org/licenses/LICENSE-2.0"
  }
}
```

#### Yarn

To generate an SBOM using Yarn, install the CycloneDX Yarn module globally or as a development dependency:

```bash
yarn add -D @cyclonedx/bom
```

Then, run:

```bash
cyclonedx-bom -o sbom.json
```

The SBOM will be generated in the current directory with the name `sbom.json`.

#### Advanced Configuration

##### Adding Metadata

You can add metadata to the SBOM using the CycloneDX Yarn module:

```json
{
  "organizationalEntity": {
    "name": "Your Organization",
    "url": "https://your-org.com",
    "contact": {
      "name": "Contact Name",
      "email": "contact@your-org.com"
    }
  }
}
```

##### Adding License Information

You can specify license information for your project:

```json
{
  "license": {
    "name": "Apache License 2.0",
    "url": "https://www.apache.org/licenses/LICENSE-2.0"
  }
}
```

### Python

#### PIP

Install the CycloneDX Python module:

```bash
pip install cyclonedx-bom
```

##### Generating the SBOM

To generate the SBOM, run:

```bash
cyclonedx-py -o sbom.json
```

The SBOM will be generated in the current directory with the name `sbom.json`.

#### Advanced Configuration

##### Adding Metadata

You can add metadata to the SBOM using the CycloneDX Python module:

```python
metadata = {
    "organizationalEntity": {
        "name": "Your Organization",
        "url": "https://your-org.com",
        "contact": {
            "name": "Contact Name",
            "email": "contact@your-org.com"
        }
    }
}
```

##### Adding License Information

You can specify license information for your project:

```python
license_info = {
    "license": {
        "name": "Apache License 2.0",
        "url": "https://www.apache.org/licenses/LICENSE-2.0"
    }
}
```

### Go

To generate an SBOM for Go projects, use the CycloneDX Go module:

Install the module:

```bash
go install github.com/CycloneDX/cyclonedx-gomod/cmd/cyclonedx-gomod@latest
```

Generate the SBOM:

```bash
cyclonedx-gomod sbom -o sbom.json
```

The SBOM will be generated in the current directory with the name `sbom.json`.

#### Advanced Configuration

##### Adding Metadata

You can add metadata to the SBOM using the CycloneDX Go module:

```go
metadata := Metadata{
    OrganizationalEntity: OrganizationalEntity{
        Name: "Your Organization",
        URL: "https://your-org.com",
        Contact: Contact{
            Name: "Contact Name",
            Email: "contact@your-org.com",
        },
    },
}
```

##### Adding License Information

You can specify license information for your project:

```go
license := License{
    Name: "Apache License 2.0",
    URL: "https://www.apache.org/licenses/LICENSE-2.0",
}
```

## SCA Tools - CI/CD Integration

### Syft

Syft is a CLI tool and library for generating SBOMs from container images and filesystems.

Install Syft:

```bash
curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh
```

Generate an SBOM:

```bash
syft <image>:<tag> -o cyclonedx-json > sbom.json
```

### CDXGen

CDXGen is a tool to generate CycloneDX SBOMs for various ecosystems.

Install CDXGen:

```bash
npm install -g @cyclonedx/cdxgen
```

Generate an SBOM:

```bash
cdxgen -o sbom.json
```

## References

- [CycloneDX Gradle Plugin Documentation](https://github.com/CycloneDX/cyclonedx-gradle-plugin)
- [CycloneDX Maven Plugin Documentation](https://github.com/CycloneDX/cyclonedx-maven-plugin)
- [CycloneDX Node.js Documentation](https://github.com/CycloneDX/cyclonedx-node-module)
- [CycloneDX Python Documentation](https://github.com/CycloneDX/cyclonedx-python)
- [CycloneDX Go Module Documentation](https://github.com/CycloneDX/cyclonedx-gomod)
- [Syft Documentation](https://github.com/anchore/syft)
- [CDXGen Documentation](https://github.com/CycloneDX/cdxgen)
